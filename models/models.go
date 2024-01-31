package models

import (
	"github.com/ethereum/go-ethereum/common"
)

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
	Funding struct {
		NextEpoch   string `json:"next_epoch"`
		FundingRate string `json:"funding_rate"`
	}
	AevoTime struct {
		Name      string `json:"name"`
		Timestamp int    `json:"timestamp"`
		Time      string `json:"time"`
		Sequence  int    `json:"sequence"`
		Block     int    `json:"block"`
	}
	AevoSignedOrder struct {
		IsBuy      bool           `json:"is_buy"`
		Instrument int            `json:"instrument"`
		LimitPrice int64          `json:"limit_price,string"`
		Amount     int64          `json:"amount,string"`
		Timestamp  int64          `json:"timestamp,string"`
		Salt       int64          `json:"salt,string"`
		Maker      common.Address `json:"maker"`
		Signature  string         `json:"signature"`
	}
)
