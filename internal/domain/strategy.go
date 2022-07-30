package domain

import (
	"main/internal/model/action"

	"github.com/adshao/go-binance/v2"
)

type Strategy interface {
	Calculate([]*binance.Kline) action.Action
}
