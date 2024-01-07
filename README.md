# aevo-go-sdk  
The community version of Go Aevo sdk  
## Examples  
### You need to add the env value from the .env file, for example  
```  
err := godotenv.Load()  
if err != nil {  
	panic(err)  
}  
```  
### Returns the new aevo client instance.   
```  
client, err := aevo.NewClient("mainnet", os.Getenv("PRIVATE_KEY")) // or "testnet"  
if err != nil {  
	panic(err)  
}  
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
## More examples in examples/public-api/example.go and examples/private-api/example.go