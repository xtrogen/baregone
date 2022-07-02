package baregone

import "time"

type TradeType string

const (
	BUY  TradeType = "BUY"
	SELL           = "SELL"
)

type BarData struct {
	date   time.Time
	close  int
	volume int
}

type BacktestContext struct {
	trades      []Position `json:"trades,omitempty"`
	capital     int        `json:"capital,omitempty"`
	profit      int        `json:"profit,omitempty"`
	totalTrades int        `json:"total_trades,omitempty"`
}
