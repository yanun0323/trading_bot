package service

import (
	"main/internal/domain"
	"main/internal/model"
)

type MockBot struct {
	symbol   string
	base     model.Balance
	currency model.Balance
	strategy Strategy
	orders   []model.MockOrder
	provider domain.KlineProvider
}

func NewMockBot(symbol string, base, currency model.Balance, strategy Strategy, provider domain.KlineProvider) *MockBot {
	return &MockBot{
		symbol:   symbol,
		base:     base,
		currency: currency,
		strategy: strategy,
		orders:   make([]model.MockOrder, 0, 20),
		provider: provider,
	}
}

func (bot *MockBot) Run() {

}

func (bot *MockBot) Stop() {

}
