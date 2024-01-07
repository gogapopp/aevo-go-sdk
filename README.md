# aevo-go-sdk
The community version of Go Aevo sdk

Examples  
Client  
```client := aevo.NewClient("https://api.aevo.xyz/")```
Returns the list of active underlying assets.  
```assets, err := client.GetAssets()  
if err != nil {  
	return err  
}```  
Returns the current index price of the given asset.  
```assetPrice, err := client.GetAssetPrice("ETH")  
if err != nil {  
	return err  
}```  