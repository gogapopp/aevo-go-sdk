package models

type AssetPrice struct {
	Timestamp int64   `json:"timestamp,string"`
	Price     float64 `json:"price,string"`
}
