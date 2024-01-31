package aevo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gogapopp/aevo-go-sdk/models"
)

// GetAssets returns the list of active underlying assets
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

// GetExpiries returns the expiry timestamps of derivatives of the given asset
func (c *Client) GetExpiries(asset string) ([]string, error) {
	url := fmt.Sprintf("%sexpiries?asset=%s", c.baseUrl, asset)
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
	expiriesString := string(body)
	expiriesString = expiriesString[1 : len(expiriesString)-1]
	expiries := strings.Split(expiriesString, ",")
	return expiries, nil
}

// GetAssetPrice returns the current index price of the given asset
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

// GetAssetPriceHistory returns the historical index price for a given asset
func (c *Client) GetAssetPriceHistory(asset string, resolution int, startTime int, endTime int) ([][]string, error) {
	url := fmt.Sprintf("%sindex-history?asset=%s&resolution=%d&start_time=%d&end_time=%d", c.baseUrl, asset, resolution, startTime, endTime)
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
	var assetPriceHistory models.AssetPriceHistory
	err = json.Unmarshal(body, &assetPriceHistory)
	if err != nil {
		return nil, err
	}
	return assetPriceHistory.AssetHistory, nil
}

// GetMarkets returns a list of instruments. If asset is not specified, the response will include all listed instruments
func (c *Client) GetMarkets(asset, instrument string) ([]byte, error) {
	url := fmt.Sprintf("%smarkets?asset=%s&instrument_type=%s", c.baseUrl, asset, instrument)
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
	return body, nil
}

// GetAssetStatistics returns the market statistics for the given asset
func (c *Client) GetAssetStatistics(asset, instrument string, endTime int) ([]byte, error) {
	url := fmt.Sprintf("%sstatistics?asset=%s&instrument_type=%s&end_time=%d", c.baseUrl, asset, instrument, endTime)
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
	return body, nil
}

// GetCoingeckoStatistics returns the perpetual statistics of all assets specifically for https://www.coingecko.com/en/exchanges/aevo
func (c *Client) GetCoingeckoStatistics() ([]models.CoingeckoStatistic, error) {
	url := fmt.Sprintf("%scoingecko-statistics", c.baseUrl)
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
	var coingeckoStat []models.CoingeckoStatistic
	err = json.Unmarshal(body, &coingeckoStat)
	if err != nil {
		return nil, err
	}
	return coingeckoStat, nil
}

// GetFunding returns the current funding rate for the instrument
func (c *Client) GetFunding(instrumentName string) (models.Funding, error) {
	url := fmt.Sprintf("%sfunding?instrument_name=%s", c.baseUrl, instrumentName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Funding{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Funding{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Funding{}, err
	}
	var funding models.Funding
	err = json.Unmarshal(body, &funding)
	if err != nil {
		return models.Funding{}, err
	}
	return funding, nil
}

// GetFundingHistory returns the funding rate history for the instrument
func (c *Client) GetFundingHistory(instrumentName string, startTime, endTime, limit int) ([]byte, error) {
	url := fmt.Sprintf("%sfunding-history?instrument_name=%s&start_time=%d&end_time=%d&limit=%d", c.baseUrl, instrumentName, startTime, endTime, limit)
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
	return body, nil
}

// GetTime returns the server time
func (c *Client) GetTime() (models.AevoTime, error) {
	url := fmt.Sprintf("%stime", c.baseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.AevoTime{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.AevoTime{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.AevoTime{}, err
	}
	var aevoTime models.AevoTime
	err = json.Unmarshal(body, &aevoTime)
	if err != nil {
		return models.AevoTime{}, err
	}
	return aevoTime, nil
}

// GetInstrumentByName returns the instrument information for the given instrument
func (c *Client) GetInstrumentByName(instrument string) ([]byte, error) {
	url := fmt.Sprintf("%sinstrument/%s", c.baseUrl, instrument)
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
	return body, nil
}
