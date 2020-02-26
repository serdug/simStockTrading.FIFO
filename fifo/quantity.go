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

// qtyNew for the added quantity, new trades.
// Note: Fees should be taken into account, so that cash is not overspent.
func (this *Asset) qtyNew(lot, fee float64) {
	this.S.Qty.I = -math.Abs(float64(this.S.Pos.I)) * lot / (this.Pxs.Tx - fee)
	this.L.Qty.I = +math.Abs(float64(this.L.Pos.I)) * lot / (this.Pxs.Tx + fee)
}

// qtyStart for the starting quantity
func (this *Asset) qtyStart(prev Asset) {
	this.S.Qty.Sta = prev.S.Qty.E
	this.L.Qty.Sta = prev.L.Qty.E
}

// qtyRemov for the quantity removed, closing all positions existing at previous 
// day's end
func (this *Asset) qtyRemov(prev Asset) {
	this.S.Qty.O = 0
	this.L.Qty.O = 0

	switch {
	case this.S.Pos.O != 0: 
		this.S.Qty.O = -prev.S.Qty.E
	
	case this.L.Pos.O != 0: 
		this.L.Qty.O = -prev.L.Qty.E
	
	default: 
		// Do nothing
	}
}

// qtyEnd for the ending quantity
func (this *Asset) qtyEnd() {
	// Variance (change) in quantity
	this.S.Qty.V = this.S.Qty.I + this.S.Qty.O
	this.L.Qty.V = this.L.Qty.I + this.L.Qty.O

	this.S.Qty.E = this.S.Qty.Sta + this.S.Qty.V
	this.L.Qty.E = this.L.Qty.Sta + this.L.Qty.V
}