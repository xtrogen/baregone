package main

import (
	"fmt"
)

func main() {

	const demoTargetPrice = 8000

	backtestArgs := Backtest{
		marketData: []BarData{},
		option: BackTestOptions{
			capital: 1000,
			debug:   true,
		},
		strategy: Strategy{
			analysePosition: func(args AnalysePositionArgs) {

			},
			onMarketTick: func(args OnMarketTickArgs) {
				if args.bar.close > demoTargetPrice {
					args.enterPosition("BUY")
				}
			},
		},
	}

	results = baregone.backtest(backtestArgs)

	fmt.Printf(results.String())
}
