package model

import (
	"main/internal/model/order"
	"main/internal/model/state"
	"time"

	"github.com/shopspring/decimal"
)

type MockOrder struct {
	OrderID uint64
	Symbol  string
	Type    order.Type
	State   state.State
	Amount  decimal.Decimal
	Price   decimal.Decimal
	CreatAt time.Time
}

func NewMockOrder(symbol string, t order.Type, amount, price decimal.Decimal) MockOrder {
	return MockOrder{
		Symbol:  symbol,
		Type:    t,
		State:   state.Complete,
		Amount:  amount,
		Price:   price,
		CreatAt: time.Now(),
	}
}
