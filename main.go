package main

import (
	"fmt"
	"os"

	"github.com/gogapopp/aevo-go-sdk/aevo"
	"github.com/gogapopp/aevo-go-sdk/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	client, err := aevo.NewClient("testnet", os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Println(err)
	}
	body, err := client.CreateAndSignOrder(models.AevoSignedOrder{
		Instrument: 933,
		IsBuy:      true,
		Amount:     "1",
		LimitPrice: "2000",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
