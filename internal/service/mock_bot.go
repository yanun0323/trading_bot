package service

import (
	"main/internal/model"

	"github.com/shopspring/decimal"
)

type MockBot struct {
	Balance  decimal.Decimal
	Strategy Strategy
	Orders   []model.MockOrder
}

func NewMockBot(balance decimal.Decimal, strategy Strategy) *MockBot {
	return &MockBot{
		Balance:  balance,
		Strategy: strategy,
		Orders:   make([]model.MockOrder, 0, 20),
	}
}

func (bot *MockBot) Run() {

}

func (bot *MockBot) Stop() {

}
