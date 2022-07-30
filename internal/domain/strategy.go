package domain

import "github.com/adshao/go-binance/v2"

// TODO: Implement Strategy
type Strategy interface {
	Calculate(*[]binance.Kline)
}
