// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

// removals for the quantity and the basis of closed position(s)
func (this *Asset) removals(prev Asset) {
	var (
		sh, ln []Pending
		qtyS, qtyL float64
		basS, basL float64
	)
	this.S.Qty.O   = 0
	this.S.Basis.O = 0

	this.L.Qty.O   = 0
	this.L.Basis.O = 0

	// Remove elements from the queue of pending positions.
	switch {
	case this.S.Pos.O != 0:
		sh, this.S.Queue = split(prev.S.Queue, this.S.Pos.O)

		for i := range sh {
			qtyS += prev.S.Queue[i].Qty
			basS += prev.S.Queue[i].Basis
		}
		this.S.Qty.O   = -qtyS
		this.S.Basis.O = -basS

	case this.L.Pos.O != 0:
		ln, this.L.Queue = split(prev.L.Queue, this.L.Pos.O)

		for j := range ln {
			qtyL += prev.L.Queue[j].Qty
			basL += prev.L.Queue[j].Basis
		}
		this.L.Qty.O   = -qtyL
		this.L.Basis.O = -basL

	default:
		// Do nothing
	}
}
