package app

import (
	"context"
	"errors"
	"fmt"
	"main/internal/model/interval"

	"github.com/adshao/go-binance/v2"
	"github.com/spf13/viper"
)

/*
symbol like "BTCUSDT"
*/
func NewApp(symbol string) error {
	if len(symbol) == 0 {
		return errors.New("empty symbol")
	}

	client := NewClient()
	// futuresClient := binance.NewFuturesClient(apiKey, secretKey)   // USDT-M Futures
	// deliveryClient := binance.NewDeliveryClient(apiKey, secretKey) // Coin-M Futures

	// InitKline(client, symbol, interval.Min15)

	return nil
}

func NewClient() *binance.Client {
	return binance.NewClient(viper.GetString("token.binance.key"), viper.GetString("token.binance.secret"))
}

func InitKline(client *binance.Client, symbol string, interval interval.Interval) {
	klines, err := client.NewKlinesService().Symbol(symbol).
		Interval(interval.String()).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, k := range klines {
		fmt.Printf("%+v\n", k)
	}

	fmt.Printf("kline count: %d - %s\n", len(klines), interval.String())
}

func InitWebSocket(symbol string, interval interval.Interval) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe(symbol, interval.String(), wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
