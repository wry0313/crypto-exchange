package ws

import "github/wry-0313/exchange/exchange"

type WebSocket struct {
	exchangeService exchange.Service
}

func NewWebSocket(exchangeService exchange.Service) *WebSocket {
	return &WebSocket{
		exchangeService: exchangeService,
	}
}