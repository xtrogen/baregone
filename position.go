package baregone

import "time"

type Position struct {
	tradeType         TradeType
	entryTime         time.Time
	exitTime          time.Time
	entryPrice        int
	profit            int
	profitAmount      int
	profitPct         int
	isOpen            bool
	virtualEntryPrice int
	virtualEntryTime  time.Time
	virtualProfit     int
}

func (p *Position) TradeType() TradeType {
	return p.tradeType
}

func (p *Position) SetTradeType(tradeType TradeType) {
	p.tradeType = tradeType
}

func (p *Position) EntryTime() time.Time {
	return p.entryTime
}

func (p *Position) SetEntryTime(entryTime time.Time) {
	p.entryTime = entryTime
}

func (p *Position) ExitTime() time.Time {
	return p.exitTime
}

func (p *Position) SetExitTime(exitTime time.Time) {
	p.exitTime = exitTime
}

func (p *Position) EntryPrice() int {
	return p.entryPrice
}

func (p *Position) SetEntryPrice(entryPrice int) {
	p.entryPrice = entryPrice
}

func (p *Position) Profit() int {
	return p.profit
}

func (p *Position) SetProfit(profit int) {
	p.profit = profit
}

func (p *Position) ProfitAmount() int {
	return p.profitAmount
}

func (p *Position) SetProfitAmount(profitAmount int) {
	p.profitAmount = profitAmount
}

func (p *Position) ProfitPct() int {
	return p.profitPct
}

func (p *Position) SetProfitPct(profitPct int) {
	p.profitPct = profitPct
}

func (p *Position) SetIsOpen(isOpen bool) {
	p.isOpen = isOpen
}

func (p *Position) IsOpen() bool {
	return p.isOpen
}
