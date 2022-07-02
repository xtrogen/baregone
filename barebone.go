package baregone

import (
	"fmt"
	"time"

	lop "github.com/samber/lo/parallel"
)

type Backtest struct {
	marketData []BarData
	option     BackTestOptions
	strategy   Strategy
}

func (args Backtest) backtest() {
	strategy := args.strategy
	isDebug := args.option.debug || false
	prices := args.marketData
	initCapital := args.option.capital || 1000
	totalPrices := len(prices)

	position := &Position{}
	currentBar := &BarData{}

	context := &BacktestContext{
		trades:  nil,
		capital: initCapital,
	}

	log := func(str string, others ...int) {
		if isDebug {
			fmt.Printf(str, others...)
		}
	}

	refreshVariables := func() {

		position = &Position{
			// tradeType:         nil,
			entryTime:         time.Time{},
			exitTime:          time.Time{},
			entryPrice:        0,
			profit:            0,
			profitAmount:      0,
			profitPct:         0,
			isOpen:            false,
			virtualEntryPrice: 0,
			virtualEntryTime:  time.Time{},
			virtualProfit:     0,
		}

		currentBar = &BarData{
			date:   time.Time{},
			close:  0,
			volume: 0,
		}

	}

	recordPosition := func() {
		tradeType := position.tradeType

		// profitToSave, profitPercentage
		profitToSave := 0
		profitPercentage := 0
		closePrice := currentBar.close
		entryPrice := position.entryPrice
		if tradeType == "SELL" {
			profitPercentage = GetPercentageGain(closePrice, entryPrice)
			profitToSave = entryPrice - closePrice
		} else {
			profitPercentage = GetPercentageGain(entryPrice, closePrice)
			profitToSave = closePrice - entryPrice
		}

		log("profitPercentage     ------------>", profitPercentage)
		log("profitToSave         ------------>", profitToSave)
		position.SetProfit(profitToSave)
		position.SetProfitPct(profitPercentage)
	}

	exitPosition := func() {
	}

	enterPosition := func(tradeType TradeType) {
	}

	finishTrading := func() BacktestContext {
		if position.isOpen {
			exitPosition()
		}

		return BacktestContext{
			trades:  context.trades,
			capital: context.capital,
			profit: lop.SumBy(context.trades, func(trade Position) int {
				return trade.profit
			}),
			totalTrades: len(context.trades),
		}
	}

	for _, price := range prices {
		currentBar := price
		isOpen := position.isOpen

		if isOpen {
			recordPosition()
			strategy.analysePosition(AnalysePositionArgs{bar: currentBar, position: *position, exitPosition: exitPosition})
		} else {
			strategy.onMarketTick(OnMarketTickArgs{bar: currentBar, enterPosition: enterPosition})
		}
	}
	// ... logic here

	// return nil
}
