package analyse

import "github.com/xtrogen/baregone"

type Analyse struct {
	bar       baregone.BarData
	positions baregone.Position
}

type AnalysePosition interface {
	exitPosition(tradeType baregone.TradeType) int
}
