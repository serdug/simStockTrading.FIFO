// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

// Asset for the results of simulated trading
type Asset struct {
	// Bar ID, e.g. date/time stamp, in any convenient format, not unique values are allowed
	Bar       string
	
	// Underlying asset's prices
	Pxs       Prices

	// Results for each side of bet are considered separately
	// The Short side
	S         Position
	// The Long side
	L         Position
	
	// Cumulative return
	CumReturn float64

	// A flag indicating that an exit trade is blocked due to a possible loss
	Block     bool

	// Net Asset Value
	NAV       float64

	// Maximal NAV attained
	MaxNAV    float64
	
	Drawdown  float64

	// Worst (maximum) drawdown
	WDD       float64

	// Trade entry counter
	EntryN    int
	
	// Trade exit counter
	ExitN     int
}

// Position for calculated results - comprehensive, spreadsheet-like
type Position struct {
	// The size of position (In-Out-Ending)
	Pos    IOE

	// The quantity of shares in position
	Qty    Tally
	
	// The basis of position
	Basis  Tally
	
	// The net market value of position
	Val    float64
	
	// FIFO queue
	Queue  []Pending
	
	// Net cash flow, net proceeds
	NetCF  Tally

	// The results of trading
	Result TReturn
}

// IOE for signals (In-Out-End)
type IOE struct {
	// The size of a new bet (an addition to the position), entry signal
	I int
	
	// The size of removed position, exit signal
	O int
	
	// The resulting size of position
	E int
}

// Prices for Close (Last) and Trade prices
type Prices struct {
	// Close (last) price
	Cl float64
	
	// Trade price
	Tx float64
}

// Tally describes changing amounts or values
type Tally struct {
	// The starting amount
	Sta float64 

	// The added (In) amount
	I   float64 

	// The removed (Out) amount
	O   float64 

	// The Variance (change) of amount
	V   float64 

	// The Ending amount
	E   float64 
}

// TReturn describes trading returns
type TReturn struct {
	// The Unrealized (Mark-to-Market) result
	Unr    float64 

	// The change of unrealized (Mark-to-Market) result
	UnrChg float64 

	// The Realized result
	Rzd    float64 

	// The Total (Realized + Unrealized) result
	Tot    float64 
}

// argsFIFO holds arguments for the fifo() function.
type argsFIFO struct {
	// Prices and signals
	Sigs     []Trades
	
	// Cash base (Cash Allocated for Trading) - assets initially allocated for 
	// the trading program
	Cashbase float64
	
	// Cash limit of exposure per position
	Lim      float64
	
	// Broker's commission
	Fee      float64
}

// fifo calculates results of model trade on the basis of signals.
// Note: no error handling.
func fifo(q argsFIFO) (values []Asset, err error) {
	var (
		this Asset
	)
	// Allocate space for a slice of trade results
	values = make([]Asset, len(q.Sigs))

	for i, signals := range q.Sigs {
		// Put underlying asset's prices into the Asset object
		this.prices(signals)

		if i == 0 {
			// Initial signals (bar 1)
			this.iniSignals(signals)

			this.qtyStart(Asset{})
			this.basisStart(Asset{})
			this.additions(Asset{}, q)
			this.qtyEnd()
			this.basisEnd()

			// Starting Net Asset Value
			this.NAV = q.Cashbase
		}

		if i > 0 {
			// Put the signals into the Asset object
			this.signals(signals, values[i-1])

			// Ending position size
			this.posEnd(values[i-1])

			// Starting quantity
			this.qtyStart(values[i-1])

			// Starting basis
			this.basisStart(values[i-1])

			// New trades, opened positions
			this.additions(values[i-1], q)

			// Closed positions
			this.removals(values[i-1])

			// Ending quantity
			this.qtyEnd()

			// Ending basis
			this.basisEnd()

			// Mark to Market
			this.mtm()

			// Net proceeds from position removal, closing an exact number of 
			// positions
			this.cfRemov(q.Fee)

			// Returns
			this.returns(values[i-1])

			// Net Asset Value
			this.assets(q.Cashbase)

			// Peak Net Asset Value
			this.maxAssets(values[i-1])

			// Drawdown
			this.drawdown(q.Cashbase)

			// Worst (maximum) drawdown
			this.wdd(values[i-1])

			// Number of trades
			// The counter of additions (initiated trades)
			this.countEntries()

			// The counter of removals (completed trades)
			this.countExits()
		}
		values[i] = this
	}
	return
}
