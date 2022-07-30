package service

import (
	"main/internal/model/action"

	"github.com/adshao/go-binance/v2"
)

type StrategyService struct {
}

func NewStrategyService() *StrategyService {
	return &StrategyService{}
}

func (svc *StrategyService) Calculate(data []*binance.Kline) action.Action {
	// TODO: Implement Strategy
	return action.None
}
