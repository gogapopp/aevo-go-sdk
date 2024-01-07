package publicapi

import (
	"fmt"

	"github.com/gogapopp/aevo-go-sdk/aevo"
)

func ExamplePublicAPI() {
	client := aevo.NewClient("https://api.aevo.xyz/")
	assets, err := client.GetAssets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(assets)
	expiries, err := client.GetExpiries("ETH")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(expiries)
	assetPrice, err := client.GetAssetPrice("ETH")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(assetPrice)
	assetPriceHistory, err := client.GetAssetPriceHistory("ETH", 30, 1672531200000000000, 1675036800000000000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(assetPriceHistory)
	markets, err := client.GetMarkets("BTC", "PERPETUAL")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(markets))
	assetStat, err := client.GetAssetStatistics("ETH", "OPTION", 1675)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(assetStat))
}
