package aevo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gogapopp/aevo-go-sdk/models"
)

func generateSalt() int {
	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Intn(900000) + 100000
}

func getHeaders(path, method, body string) map[string]string {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	message := fmt.Sprintf("%s,%s,%s,%s,%s", apiKey, timestamp, strings.ToUpper(method), path, body)
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))

	headers := map[string]string{
		"AEVO-TIMESTAMP": timestamp,
		"AEVO-SIGNATURE": signature,
		"AEVO-KEY":       apiKey,
	}

	return headers
}

func (c *Client) CreateAndSignOrder(order models.AevoSignedOrder) ([]byte, error) {
	// construct the order
	order.Timestamp = time.Now().Unix()
	order.Salt = int64(generateSalt())
	// check signing key
	if c.signingKey == nil {
		return nil, fmt.Errorf("signing key is not provided")
	}
	order.Maker = crypto.PubkeyToAddress(c.signingKey.PublicKey)

	types := apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			// {Name: "salt", Type: "string"},
			// {Name: "verifyingContract", Type: "address"},
		},
		"Order": {
			{Name: "maker", Type: "address"},
			{Name: "isBuy", Type: "bool"},
			{Name: "limitPrice", Type: "uint256"},
			{Name: "amount", Type: "uint256"},
			{Name: "salt", Type: "uint256"},
			{Name: "instrument", Type: "uint256"},
			{Name: "timestamp", Type: "uint256"},
		},
	}
	// privateKeys, err := crypto.HexToECDSA(os.Getenv("SIGNING_KEY"))
	// if err != nil {
	// 	return nil, err
	// }
	// name='Aevo Testnet', version='1', chainId=11155111
	// name='Aevo Mainnet', version='1', chainId=1
	chainID := math.NewHexOrDecimal256(11155111)
	domain := apitypes.TypedDataDomain{
		Name:    "Aevo Testnet",
		ChainId: chainID,
		Version: "1",
		// Salt:    hexutil.Encode(crypto.Keccak256([]byte(time.Now().String()))),
		// VerifyingContract: crypto.PubkeyToAddress(privateKeys.PublicKey).Hex(),
	}
	message := map[string]interface{}{
		"isBuy":      order.IsBuy,
		"instrument": big.NewInt(int64(order.Instrument)),
		"limitPrice": big.NewInt(int64(order.LimitPrice)),
		"amount":     big.NewInt(int64(order.Amount)),
		"timestamp":  big.NewInt(int64(order.Timestamp)),
		"salt":       big.NewInt(int64(order.Salt)),
		"maker":      order.Maker.String(),
	}

	typedData := apitypes.TypedData{
		Types:       types,
		PrimaryType: "Order",
		Domain:      domain,
		Message:     message,
	}
	// hash structures
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	// sign with key
	signingKey := os.Getenv("SIGNING_KEY")
	key, err := crypto.HexToECDSA(signingKey)
	if err != nil {
		return nil, err
	}
	// construct
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	dataHash := crypto.Keccak256(rawData)
	// hash := common.BytesToHash(hashBytes)

	signature, err := crypto.Sign(dataHash, key)
	if err != nil {
		return nil, err
	}

	// if signature[64] < 27 {
	// 	signature[64] += 27
	// }

	fmt.Printf("Signature: %s\n", hexutil.Encode(signature))
	// order.Signature = fmt.Sprintf("0x%s", hexutil.Encode(signature))
	order.Signature = hexutil.Encode(signature)

	// TODO: EIP712 signature gen

	// create the order
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	fmt.Println(order)
	url := fmt.Sprintf("%s/orders", c.baseUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(orderJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	headers := getHeaders("/orders", "POST", string(orderJSON))
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) CancelOrder(orderID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.baseUrl, orderID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	headers := getHeaders("/"+orderID, "DELETE", "")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) GetOrders() ([]byte, error) {
	url := fmt.Sprintf("%sorders", c.baseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	headers := getHeaders("/orders", "GET", "")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
