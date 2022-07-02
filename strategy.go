package baregone

type BackTestOptions struct {
	capital int
	debug   bool
}

type BackTestParams struct {
	symbol     string    `json:"symbol,omitempty"`
	marketData []BarData `json:"market_data,omitempty"`
}

type EnterPosition func(tradeType TradeType)
type ExitPosition func()

type AnalysePositionArgs struct {
	bar          BarData
	position     Position
	exitPosition ExitPosition
}

type OnMarketTickArgs struct {
	bar           BarData
	enterPosition EnterPosition
}

type AnalysePosition func(args AnalysePositionArgs)
type OnMarketTick func(args OnMarketTickArgs)

// move to strategy
type Strategy struct {
	analysePosition AnalysePosition
	onMarketTick    OnMarketTick
}
