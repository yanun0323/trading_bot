package service

import (
	"main/internal/domain"
	"main/internal/model"
)

type MockBot struct {
	symbol   string
	base     *model.Balance
	currency *model.Balance
	orders   []model.MockOrder
	strategy domain.Strategy
	provider domain.KlineProvider
}

func NewMockBot(symbol string, balance string, strategy domain.Strategy, provider domain.KlineProvider) *MockBot {
	return &MockBot{
		symbol:   symbol,
		base:     model.NewBalance(symbol[3:len(symbol)-1], balance, "0"),
		currency: model.NewBalance(symbol[:2], "0", "0"),
		strategy: strategy,
		orders:   make([]model.MockOrder, 0, 20),
		provider: provider,
	}
}

func (bot *MockBot) Run() {

}

func (bot *MockBot) Stop() {

}
