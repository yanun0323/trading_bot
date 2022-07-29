package app

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/spf13/viper"
)

func NewApp() {
	client := binance.NewClient(viper.GetString("token.binance.key"), viper.GetString("token.binance.secret"))
	fmt.Println(client)
	// futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	// deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures
}
