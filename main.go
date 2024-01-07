package main

import (
	"fmt"

	"github.com/gogapopp/aevo-go-sdk/aevo"
)

func main() {
	client := aevo.NewClient("https://api.aevo.xyz/")
	assets, err := client.GetAssets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(assets)
	assetPrice, err := client.GetAssetPrice("ETH")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(assetPrice)
}
