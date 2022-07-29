package domain

import (
	"main/internal/model/interval"

	"github.com/adshao/go-binance/v2"
)

type KlineProvider interface {
	GetKline(string, interval.Interval) []*binance.Kline
}
