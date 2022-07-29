package service

import (
	"main/internal/model/interval"

	"github.com/adshao/go-binance/v2"
)

type Provider struct {
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) GetKline(symbol string, interval interval.Interval) []*binance.Kline {
	// TODO: Implement Provider
	return nil
}
