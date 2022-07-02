package baregone

import (
	"time"
)

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

/**
 * Get Profit percentage gained
 * https://www.investopedia.com/ask/answers/how-do-you-calculate-percentage-gain-or-loss-investment/
 * @param startPrice
 * @param endPrice
 */
func GetPercentageGain(startPrice int, endPrice int) int {
	if startPrice <= 0 {
		return 0
	}

	if endPrice <= 0 {
		return 0
	}

	return (endPrice - startPrice) / startPrice * 100
}

/**
 * GetTotalProfitAmount, from start and end
 * @param start
 * @param end
 * @param capital
 */
func GetTotalProfitAmount(start int, end int, capital int) int {
	// TODO remove
	if start <= 0 {
		return 0
	}

	if end <= 0 {
		return 0
	}

	profit := end - start
	return (profit / start) * capital
}
