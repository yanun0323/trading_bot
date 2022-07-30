package util

import (
	"context"
	"errors"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func ConnectAccount(client *binance.Client) (*binance.Account, error) {
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("new get account service, %w", err)
	}
	return res, nil
}

func GetBalance(client *binance.Client, symbols ...string) (map[string]binance.Balance, error) {
	account, err := ConnectAccount(client)
	if err != nil {
		return nil, fmt.Errorf("connect account, %w", err)
	}
	if len(symbols) == 0 {
		return nil, errors.New("empty symbol")
	}

	set := make(map[string]interface{}, 0)
	for _, symbol := range symbols {
		set[symbol] = nil
	}
	result := make(map[string]binance.Balance)
	for _, balance := range account.Balances {
		if _, exist := set[balance.Asset]; exist {
			result[balance.Asset] = balance
		}
	}
	return result, nil
}
