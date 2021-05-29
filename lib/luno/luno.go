package luno

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OrderStatus int

const (
	Placed OrderStatus = 1
	Traded OrderStatus = 2
)

// Client simulates a Luno API trading client by estimating the state of placed order against the real live market.
type Client struct{}

// PlacePostOnly simulates placing a post only order on the Client exchange.
func (l Client) PlacePostOnly(market string, isBuy bool, amount float64, price float64, extID string) error {
	return nil
}

// Cancel simulates cancelling an order on the Client exchange.
func (l Client) Cancel(extID string) error {
	return nil
}

// Ticker returns the actual market ticker.
func (l Client) Ticker(market string) (Ticker, error) {
	r, err := http.Get("https://api.luno.com/api/1/ticker?pair=" + market)
	if err != nil {
		return Ticker{}, err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return Ticker{}, err
	}

	var t Ticker
	err = json.Unmarshal(b, &t)
	if err != nil {
		return Ticker{}, err
	}

	return t, nil
}

// RecentTrades returns a list of actual recent trades.
func (l Client) RecentTrades(market string, since int64) ([]Trade, error) {
	r, err := http.Get("https://api.luno.com/api/1/trades?pair=" + market + "&since=" + fmt.Sprint(since))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	res := struct {
		Trades []Trade
	}{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res.Trades, nil
}

// GetStatus estimates the state of the given order against recent live trades.
func (l Client) GetStatus(market string, price float64, isBuy bool, timeUnixMilli int64) (OrderStatus, error) {
	tl, err := l.RecentTrades(market, timeUnixMilli)
	if err != nil {
		return 0, err
	}

	for _, trade := range tl {
		if isBuy == trade.IsBuy {
			continue
		}
		if isBuy && trade.Price <= price {
			return Traded, nil
		} else if !isBuy && trade.Price >= price {
			return Traded, nil
		}
	}

	return Placed, nil
}

type Ticker struct {
	Ask       float64 `json:"ask,string"`
	Bid       float64 `json:"bid,string"`
	LastTrade float64 `json:"last_trade,string"`
}

type Trade struct {
	IsBuy     bool    `json:"is_buy"`
	Price     float64 `json:"price,string"`
	Timestamp int64   `json:"timestamp"`
}
