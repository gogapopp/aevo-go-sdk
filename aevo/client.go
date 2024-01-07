package aevo

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Client struct {
	baseUrl    string
	chainType  string
	signingKey *ecdsa.PrivateKey
}

const (
	testnetUrl = "https://api-testnet.aevo.xyz/"
	mainnetUrl = "https://api.aevo.xyz/"
)

func NewClient(chainType, privateKeyString string) (*Client, error) {
	baseUrl := testnetUrl
	if chainType == "mainnet" {
		baseUrl = mainnetUrl
	}
	// парсим приватный ключ
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return &Client{}, err
	}
	return &Client{
		baseUrl:    baseUrl,
		chainType:  chainType,
		signingKey: privateKey,
	}, nil
}
