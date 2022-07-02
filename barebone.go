package baregone

import "time"

type Backtest struct {
	marketData []BarData
	option     BackTestOptions
	strategy   Strategy
}

func (args Backtest) backtest() {
	strategy := args.strategy
	isDebug := args.option.debug || false
	prices := args.marketData
	capital := args.option.capital || 1000
	totalPrices := len(prices)

	position := &Position{}
	currentBar := &BarData{}

	refreshVariables := func() {
		currentBar := &BarData{
			date:   time.Now(),
			close:  0,
			volume: 0,
		}

	}

	recordPosition := func() {
	}

	finishTrading := func() {
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
