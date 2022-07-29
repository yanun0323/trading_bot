package model

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type Balance struct {
	Asset string
	Free  decimal.Decimal
	Lock  decimal.Decimal
}

func NewBalance(asset string, free, lock string) *Balance {
	if len(asset) == 0 {
		return nil
	}

	f, err := decimal.NewFromString(free)
	if err != nil {
		return nil
	}

	l, err := decimal.NewFromString(lock)
	if err != nil {
		return nil
	}

	return &Balance{
		Asset: strings.ToUpper(asset),
		Free:  f,
		Lock:  l,
	}
}

func (b *Balance) TakeOut(amount decimal.Decimal) (decimal.Decimal, error) {
	if amount.IsNegative() {
		return decimal.Zero, errors.New("amount must be positive")
	}
	if amount.GreaterThan(b.Free) {
		return decimal.Zero, errors.New("you don't have enough balance in " + b.Asset)
	}
	b.Free = b.Free.Sub(amount)
	return amount, nil
}

func (b *Balance) PutIn(amount decimal.Decimal) error {
	if amount.IsNegative() {
		return errors.New("amount must be positive")
	}
	b.Free = b.Free.Add(amount)
	return nil
}
