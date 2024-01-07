package aevo

type Client struct {
	baseUrl   string
	chainType string
}

const (
	testnetUrl = "https://api-testnet.aevo.xyz/"
	mainnetUrl = "https://api.aevo.xyz/"
)

func NewClient(chainType string) *Client {
	// testnet by default
	baseUrl := testnetUrl
	if chainType == "mainnet" {
		baseUrl = mainnetUrl
	} else if chainType == "testnet" {
		baseUrl = testnetUrl
	}
	return &Client{
		baseUrl:   baseUrl,
		chainType: chainType,
	}
}
