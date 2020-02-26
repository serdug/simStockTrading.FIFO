// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

// additions adds newly open positions to the queue. Positions are split up 
// into elements (size = 1).
// Note: Fees are taken into account, so that cash is not overspent and the 
// exposure limit be used in full.
func (this *Asset) additions(prev Asset, q argsFIFO) {
	var (
		sh, ln []Pending
	)

	this.qtyNew(q.Lim, q.Fee)
	this.cfNew(q.Fee)
	this.basisNew()

	// Add elements the queue of pending positions.
	switch {
	case this.S.Pos.I != 0:
		sh = make([]Pending, this.S.Pos.I)

		// The net cost price received from selling one element of short position 
		// (size = 1).
		price2receive := this.Pxs.Tx - q.Fee

		for i := range sh {
			// Selling to open a short position. Price received, fee subtracted.
			sh[i] = Pending{
				// Note: Sign convention. The short stock has a negative quantity.
				Qty:   -q.Lim / price2receive,
				// Note: Sign convention. All costs are positive.
				Cost:  price2receive,
				// Note: Sign convention. SHORT ==> positive proceeds, negative basis.
				Basis: -q.Lim,
			}
		}

	case this.L.Pos.I != 0:
		ln = make([]Pending, this.L.Pos.I)

		// The net cost price paid for buying one element of long position 
		// (size = 1).
		price2pay := this.Pxs.Tx + q.Fee

		for j := range ln {
			// Buying to open a long position. Price paid, fee added.
			ln[j] = Pending{
				// Note: Sign convention. The long stock has a positive quantity.
				Qty:   +q.Lim / price2pay,
				// Note: Sign convention. All costs are positive.
				Cost:  price2pay,
				// Note: Sign convention. LONG ==> negative proceeds, positive basis.
				Basis: q.Lim,
			}
		}

	default:
		// Do nothing
	}
	this.S.Queue = queueAdd(prev.S.Queue, sh)
	this.L.Queue = queueAdd(prev.L.Queue, ln)
}

