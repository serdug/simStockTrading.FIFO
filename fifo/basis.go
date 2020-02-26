// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

// basisStart for the starting basis
func (this *Asset) basisStart(prev Asset) {
	this.S.Basis.Sta = prev.S.Basis.E
	this.L.Basis.Sta = prev.L.Basis.E
}

// basisNew adds the basis of newly opened trades
func (this *Asset) basisNew() {
	this.S.Basis.I = -this.S.NetCF.I
	this.L.Basis.I = -this.L.NetCF.I
}

// basisEnd for the resulting basis
func (this *Asset) basisEnd() {
	// Variance (change) in the basis
	this.S.Basis.V = this.S.Basis.I + this.S.Basis.O
	this.L.Basis.V = this.L.Basis.I + this.L.Basis.O

	this.S.Basis.E = this.S.Basis.Sta + this.S.Basis.V
	this.L.Basis.E = this.L.Basis.Sta + this.L.Basis.V
}
