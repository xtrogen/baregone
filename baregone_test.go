package baregone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"

	lop "github.com/samber/lo"
)

type MD struct {
	Price string
	Time  string
}
type JsonFile struct {
	Date       int
	MarketData []MD
}

func TestBacktest(t *testing.T) {

	file, _ := ioutil.ReadFile("BTCUSDT.json")

	data := JsonFile{}

	_ = json.Unmarshal([]byte(file), &data)

	// log.Print("original", data.MarketData[1])

	marketData := lop.Map(data.MarketData, func(md MD, i int) BarData {

		dateUnix, _ := strconv.ParseInt(md.Time, 10, 64)
		price, _ := strconv.ParseInt(md.Price, 10, 64)

		return BarData{
			date:   time.Unix(dateUnix/1000, 0),
			close:  int(price),
			volume: 0,
		}
	})

	// log.Print("parsed", marketData)

	const demoTargetPrice = 8100

	backtestArgs := BacktestArgs{
		marketData: marketData,
		option: BackTestOptions{
			capital: 1000,
			debug:   true,
		},
		strategy: Strategy{
			analysePosition: func(args AnalysePositionArgs) {
				// log.Print("analysePosition ----------------", args.bar.close)

				log.Print("analysePosition ----------------", args.position.tradeType)

				if args.position.profit <= 0 {
					log.Print("----------------closing position ----------------")
					args.exitPosition()
				}

			},
			onMarketTick: func(args OnMarketTickArgs) {
				if args.bar.close >= demoTargetPrice {
					args.enterPosition("BUY")
				}
				// log.Print("currentBar ----------------", args.bar.close)
			},
		},
	}

	results := Backtest(backtestArgs)

	fmt.Println("totalTrades ---------------------------", results.totalTrades)
	fmt.Println("profit --------------------------------", results.profit)
	fmt.Println("profit --------------------------------", results.trades)
}
