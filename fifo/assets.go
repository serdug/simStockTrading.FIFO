// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

const significance float64 = 0.0001

// assets for net asset values
func (this *Asset) assets(cashbase float64) {
	this.NAV = cashbase + this.CumReturn
}

// maxAssets for peak asset values
func (this *Asset) maxAssets(prev Asset) {
	this.MaxNAV = prev.MaxNAV

	if this.NAV > prev.MaxNAV {
		this.MaxNAV = this.NAV	
	}
}

// drawdown calculates drawdowns.
// Note: this is the ratio of a peak-to-trough decline to the initialy allocated cash.
func (this *Asset) drawdown(cashbase float64) {
	this.Drawdown = 0

	dd := (this.MaxNAV - this.NAV) / cashbase

	if dd > significance {
		this.Drawdown = dd
	}
}

// wdd for worst drawdown values
func (this *Asset) wdd(prev Asset) {
	this.WDD = prev.WDD
	if this.Drawdown > prev.WDD {
		this.WDD = this.Drawdown	
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}