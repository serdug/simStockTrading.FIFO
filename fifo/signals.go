// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

import (
	"fmt"
	"math"
	"strconv"
)

const fairSize int = 9999

// iniSignals puts initial signals into the Asset object.
// No parsing errors are taken into account.
func (this *Asset) iniSignals(sigs Trades) {
	this.Bar = sigs.Dt

	this.S.Pos.I = 0
	this.S.Pos.O = 0
	this.S.Pos.E = 0
	this.L.Pos.I = 0
	this.L.Pos.O = 0
	this.L.Pos.E = 0

	// Note: No parsing errors taken into account
	position, err := strconv.Atoi(sigs.SL)
	if err != nil {
		msg := "WARNING! '" + sigs.SL + "' read as '" + strconv.Itoa(position) + "'"
		warning(msg, err)
	}

	// Note: Sign convention. The sizes of all positions are positive.
	size := int(math.Abs(float64(position)))
	sign := sign(position)
	
	switch {
	case sign < 0:
		this.S.Pos.I = size
		this.S.Pos.E = size

	case sign > 0:
		this.L.Pos.I = size
		this.L.Pos.E = size

	default:
		// Do nothing
	}
}

// signals puts signals into the Asset object.
// No parsing errors are taken into account.
func (this *Asset) signals(sigs Trades, prev Asset) {
	this.Bar = sigs.Dt

	this.S.Pos.I = 0
	this.S.Pos.O = 0
	this.L.Pos.I = 0
	this.L.Pos.O = 0

	// Note: No parsing errors taken into account
	position, err := strconv.Atoi(sigs.SL)
	if err != nil {
		msg := "WARNING! '" + sigs.SL + "' read as '" + strconv.Itoa(position) + "'"
		warning(msg, err)
	}

	// Note: Sign convention. The sizes of all positions are positive.
	size := int(math.Abs(float64(position)))
	sign := sign(position)

	// Flip over
	short2long := sign > 0 && prev.S.Pos.E != 0
	long2short := sign < 0 && prev.L.Pos.E != 0

	// No flip-over, size goes up, position grows, increases in size
	shortGrow := sign < 0 && size > prev.S.Pos.E
	longGrow  := sign > 0 && size > prev.L.Pos.E

	// No flip-over, size goes down, position shrinks, declines in size
	shortDecl := sign < 0 && size < prev.S.Pos.E
	longDecl  := sign > 0 && size < prev.L.Pos.E

	// Position goes down to 0
	short2zero := size == 0 && prev.S.Pos.E != 0
	long2zero  := size == 0 && prev.L.Pos.E != 0

	switch {
		case short2long:
			// Exit short
			this.S.Pos.O = prev.S.Pos.E
			// Enter long
			this.L.Pos.I = size

		case long2short:
			// Exit long
			this.L.Pos.O = prev.L.Pos.E
			// Enter short
			this.S.Pos.I = size

		case shortGrow:
			// Increase the short
			this.S.Pos.I = size - prev.S.Pos.E

		case shortDecl:
			// Reduce the short
			this.S.Pos.O = prev.S.Pos.E - size

		case longGrow:
			// Increase the long
			this.L.Pos.I = size - prev.L.Pos.E

		case longDecl:
			// Reduce the long
			this.L.Pos.O = prev.L.Pos.E - size

		case short2zero:
			// Exit short
			this.S.Pos.O = prev.S.Pos.E

		case long2zero:
			// Exit long
			this.L.Pos.O = prev.L.Pos.E

		default:
			// Do nothing
		}
}

// posEnd for the ending position size.
func (this *Asset) posEnd(prev Asset) {
	// Note: Sign convention. The sizes of all positions are positive.
	this.S.Pos.E = prev.S.Pos.E + this.S.Pos.I - this.S.Pos.O
	this.L.Pos.E = prev.L.Pos.E + this.L.Pos.I - this.L.Pos.O

	// Unreasonably big position sizes are not expected.
	if this.S.Pos.E + this.L.Pos.E > fairSize {
		excl := "How about trying a position size smaller than " + strconv.Itoa(fairSize + 1) + " with a higher limit per position?"
		fmt.Printf("%s\n", excl)
	}
}

// sign returns a sign of an integer argument
func sign(x int) int {
	if x == 0 {
		return 0
	} else {
		return x / int(math.Abs(float64(x)))
	}
}
