package publicapi

import (
	"fmt"
	"os"
	"time"

	"github.com/gogapopp/aevo-go-sdk/aevo"
	"github.com/joho/godotenv"
)

func ExamplePublicAPI() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	client, err := aevo.NewClient("mainnet", os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Println(err)
	}

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

	// markets type is []byte
	markets, err := client.GetMarkets("BTC", "PERPETUAL")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(markets))

	// assetStat type is []byte
	assetStat, err := client.GetAssetStatistics("ETH", "OPTION", 1675)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(assetStat))

	coingeckoStat, err := client.GetCoingeckoStatistics()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(coingeckoStat)

	time.Sleep(time.Second * 1)

	funding, err := client.GetFunding("ETH-PERP")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(funding)

	// fundingHistory type is []byte
	fundingHistory, err := client.GetFundingHistory("ETH-PERP", 1672531200000000000, 1677036800000000000, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fundingHistory))

	aevoTime, err := client.GetTime()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aevoTime)

	// getInstrument type is []byte
	getInstrument, err := client.GetInstrumentByName("ETH-PERP")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(getInstrument))
}
