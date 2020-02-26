// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

import (
	"fmt"
	"strconv"
)

// prices puts underlying asset's Close (last) and Trade prices into the Asset 
// object.
// No parsing errors are taken into account.
func (this *Asset) prices(signals Trades) {
	closePx, errCl := strconv.ParseFloat(signals.Cl, 64)
	if errCl != nil {
		msgCl := "WARNING! '" + signals.Cl + "' read as '" + fmt.Sprintf("%f", closePx) + "'"
		warning(msgCl, errCl)
	}
	this.Pxs.Cl = closePx

	tradePx, errTx := strconv.ParseFloat(signals.Tx, 64)
	if errTx != nil {
		msgTx := "WARNING! '" + signals.Tx + "' read as '" + fmt.Sprintf("%f", tradePx) + "'"
		warning(msgTx, errTx)
	}
	this.Pxs.Tx = tradePx
}
