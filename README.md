# aevo-go-sdk
The community version of Go Aevo sdk
## Examples
### Returns the new aevo client instance.   
```
client := aevo.NewClient("mainnet") // or "testnet"
```  
### Returns the list of active underlying assets.  
```
assets, err := client.GetAssets()  
if err != nil {  
	panic(err)  
}
```  
### Returns the expiry timestamps of derivatives of the given asset.  
```  
expiries, err := client.GetExpiries("ETH")  
if err != nil {  
	panic(err)  
}  
```  
### Returns the current index price of the given asset.  
```
assetPrice, err := client.GetAssetPrice("ETH")  
if err != nil {  
	panic(err)  
}
```  
### Returns the historical index price for a given asset.  
```  
assetPriceHistory, err := client.GetAssetPriceHistory("ETH", 30, 1672531200000000000, 1675036800000000000)  
if err != nil {  
	panic(err)  
}  
```  
### Returns a list of instruments. If asset is not specified, the response will include all listed instruments.  
```  
// markets type is []byte
markets, err := client.GetMarkets("BTC", "PERPETUAL")  
if err != nil {  
	panic(err)  
}  
```  
### Returns the market statistics for the given asset.
```  
// assetStat type is []byte
assetStat, err := client.GetAssetStatistics("ETH", "OPTION", 1675036800000000000)  
if err != nil {  
	panic(err)  
}  
```  
### Returns the perpetual statistics of all assets specifically for https://www.coingecko.com/en/exchanges/aevo.
```  
coingeckoStat, err := client.GetCoingeckoStatistics()  
if err != nil {  
	panic(err)  
}  
```  

## todo:  
### public-api:  
GET /mark-history  
GET /settlement-history  
GET /orderbook  
GET /funding  
GET /funding-history  
GET /instrument/{instrument_name}  
GET /instrument/{instrument_name}/trade-history  
GET /check-referral  
POST /account/unsubscribe  
GET /time  
GET /yield-vault  
POST /swap/preview  
GET /options-history  
POST /account/email-verified  
### private-api:  
POST /register  
DELETE /api-key  
GET /api-key  
POST /api-key  
DELETE /signing-key  
GET /account  
GET /account/cancel-on-disconnect  
POST /account/cancel-on-disconnect  
POST /account/portfolio-margin  
GET /account/email-address  
POST /account/email-address  
POST /account/email-preference  
GET /account/email-preferences  
GET /account/email-verified  
GET /account/accumulated-fundings  
POST /account/update-margin  
POST /account/margin-type  
POST /account/leverage  
GET /portfolio  
POST /withdraw  
POST /strategy/initiate-withdraw  
POST /strategy/pending-transactions  
POST /transfer  
GET /orders  
POST /orders  
DELETE /orders/{order_id}  
GET /orders/{order_id}  
POST /orders/{order_id}  
DELETE /orders-all  
GET /order-history  
GET /trade-history  
GET /transaction-history  
GET /referral-rewards-history  
GET /referral-history  
GET /referral-statistics  
POST /claim-referral-rewards  
GET /mmp  
POST /mmp  
POST /reset-mmp  
DELETE /rfqs  
GET /rfqs  
POST /rfqs  
DELETE /rfqs/{block_id}  
GET /rfqs/{block_id}/quotes  
DELETE /quotes  
GET /quotes  
POST /quotes  
POST /quotes/preview  
DELETE /quotes/{quote_id}  
PUT /quotes/{quote_id}  
POST /swap  
