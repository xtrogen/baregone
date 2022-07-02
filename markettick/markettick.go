package markettick

import "github.com/xtrogen/baregone"

type MarketTick struct {
	bar baregone.BarData
}

type onMarketTick interface {
	enterPosition(tradeType baregone.TradeType) int
}
