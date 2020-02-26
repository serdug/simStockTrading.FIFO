// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

// countEntries counts trade entries, i.e. initiated trades.
func (this *Asset) countEntries() {
	if this.S.Pos.I != 0 || this.L.Pos.I != 0 {
		this.EntryN += 1
	}
}

// countExits counts trade exits, i.e. completed trades.
func (this *Asset) countExits() {
	if this.S.Pos.O != 0 || this.L.Pos.O != 0 {
		this.ExitN += 1
	}
}