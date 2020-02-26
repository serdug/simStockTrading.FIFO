// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

// mtm for market values of the asset
func (this *Asset) mtm() {
	this.S.Val = this.S.Qty.E * this.Pxs.Cl
	this.L.Val = this.L.Qty.E * this.Pxs.Cl
}

// returns for total Net Returns, realized and unrealized
func (this *Asset) returns(prev Asset) {
	this.unrealized(prev)
	this.realized()

	// Returns total
	this.S.Result.Tot = this.S.Result.Rzd + this.S.Result.Unr
	this.L.Result.Tot = this.L.Result.Rzd + this.L.Result.Unr

	// Cumulative sum of returns
	this.CumReturn += (this.S.Result.Rzd + this.L.Result.Rzd +
		this.S.Result.UnrChg + this.L.Result.UnrChg)
}

// unrealized for the Unrealized Returns
func (this *Asset) unrealized(prev Asset) {
	this.S.Result.Unr = this.S.Val - this.S.Basis.E
	this.L.Result.Unr = this.L.Val - this.L.Basis.E

	this.S.Result.UnrChg = this.S.Result.Unr - prev.S.Result.Unr
	this.L.Result.UnrChg = this.L.Result.Unr - prev.L.Result.Unr
}

// realized for total Net Realized Returns 
func (this *Asset) realized() {
	this.S.Result.Rzd = (this.S.NetCF.O + this.S.Basis.O)
	this.L.Result.Rzd = (this.L.NetCF.O + this.L.Basis.O)
}
