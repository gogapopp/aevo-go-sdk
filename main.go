package main

import (
	"fmt"

	"github.com/gogapopp/aevo-go-sdk/aevo"
)

func main() {
	client := aevo.NewClient("mainnet")
	coingeckoStat, err := client.GetCoingeckoStatistics()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(coingeckoStat)
}
