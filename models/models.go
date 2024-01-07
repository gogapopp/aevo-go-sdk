package models

type (
	AssetPrice struct {
		Timestamp int64   `json:"timestamp,string"`
		Price     float64 `json:"price,string"`
	}
	AssetPriceHistory struct {
		AssetHistory [][]string `json:"history"`
	}
	CoingeckoStatistic struct {
		TickerID                 string `json:"ticker_id"`
		BaseCurrency             string `json:"base_currency"`
		TargetCurrency           string `json:"target_currency"`
		TargetVolume             string `json:"target_volume"`
		ProductType              string `json:"product_type"`
		OpenInterest             string `json:"open_interest"`
		IndexPrice               string `json:"index_price"`
		IndexCurrency            string `json:"index_currency"`
		NextFundingRateTimestamp string `json:"next_funding_rate_timestamp"`
		FundingRate              string `json:"funding_rate"`
		ContractType             string `json:"contract_type"`
		ContractPriceCurrency    string `json:"contract_price_currency"`
	}
)
