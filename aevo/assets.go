package aevo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gogapopp/aevo-go-sdk/aevo/models"
)

func (c *Client) GetAssets() ([]string, error) {
	url := fmt.Sprintf("%sassets", c.baseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	assetsString := string(body)
	assetsString = assetsString[1 : len(assetsString)-1]
	assets := strings.Split(assetsString, ",")
	return assets, nil
}

func (c *Client) GetAssetPrice(asset string) (models.AssetPrice, error) {
	url := fmt.Sprintf("%sindex?asset=%s", c.baseUrl, asset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.AssetPrice{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.AssetPrice{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.AssetPrice{}, err
	}
	var assetPrice models.AssetPrice
	err = json.Unmarshal(body, &assetPrice)
	if err != nil {
		return models.AssetPrice{}, err
	}
	return assetPrice, nil
}
