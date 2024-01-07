package aevo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogapopp/aevo-go-sdk/models"
)

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

// func (c *Client) CreateAndSignOrder(order models.AevoSignedOrder) ([]byte, error) {
// 	// construct the order
// 	order.Timestamp = time.Now().Unix()
// 	order.Salt = int64(rand.Intn(900000) + 100000)
// 	// sign the order
// 	if c.signingKey == nil {
// 		return nil, fmt.Errorf("signing key is not provided")
// 	}
// 	privateKey := c.signingKey
// 	order.Maker = goc.PubkeyToAddress(privateKey.PublicKey)

// 	// TODO: signature gen

// 	// create the order
// 	orderJSON, err := json.Marshal(order)
// 	if err != nil {
// 		fmt.Println("4")
// 		return nil, err
// 	}
// 	fmt.Println(string(orderJSON))
// 	url := fmt.Sprintf("%s/orders", c.baseUrl)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(orderJSON))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	headers := getHeaders("/orders", "POST", string(orderJSON))
// 	for k, v := range headers {
// 		req.Header.Add(k, v)
// 	}
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, nil
// }

func (c *Client) CreateAndSignOrder(order models.AevoSignedOrder) ([]byte, error) {
	// construct the order
	order.Timestamp = time.Now().Unix()
	order.Salt = rand.Int63()
	// sign the order
	if c.signingKey == nil {
		return nil, fmt.Errorf("signing key is not provided")
	}
	order.Maker = crypto.PubkeyToAddress(c.signingKey.PublicKey)

	// TODO: EIP712 signature gen

	// // create the order
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(orderJSON))
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
