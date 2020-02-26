// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

import (
	"math"
)

// cfNew for net proceeds, or net cash flow, from new trades, opening new positions
func (this *Asset) cfNew(fee float64) {
	// Note: SHORT ==> positive proceeds; selling to open short position;
	// price received, fees subtracted
	this.S.NetCF.I = +math.Abs(this.S.Qty.I) * (this.Pxs.Tx - fee)
	// Note: LONG ===> negative proceeds; buying to open long position;
	// price paid, fees added
	this.L.NetCF.I = -math.Abs(this.L.Qty.I) * (this.Pxs.Tx + fee)
}

// cfRemov for net proceeds, or net cash flow, from position removal, closing 
// positions
func (this *Asset) cfRemov(fee float64) {
	// Note: SHORT ==> negative proceeds; buying to close short position;
	// price paid, fee added
	this.S.NetCF.O = -math.Abs(this.S.Qty.O) * (this.Pxs.Tx + fee)
	// Note: LONG ==> positive proceeds; selling to close long position;
	// price received, fees subtracted
	this.L.NetCF.O = +math.Abs(this.L.Qty.O) * (this.Pxs.Tx - fee)
}
