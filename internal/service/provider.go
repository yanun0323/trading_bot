package service

import (
	"context"
	"fmt"
	"main/internal/model/interval"

	"github.com/adshao/go-binance/v2"
)

type Provider struct {
	client *binance.Client
}

func NewProvider(client *binance.Client) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) GetKline(symbol string, interval interval.Interval) ([]*binance.Kline, error) {
	klines, err := p.client.NewKlinesService().Symbol(symbol).
		Interval(interval.String()).Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("new kline service, %w", err)
	}

	fmt.Printf("kline count: %d - %s\n", len(klines), interval.String())
	return klines, nil
}
