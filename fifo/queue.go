// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

type Pending struct {
	// The number of stocks once opened and not closed yet
	Qty   float64

	// The cost of a position once opened and not closed yet
	Cost  float64

	// The basis of a position once opened and not closed yet
	Basis float64
}

// queueAdd adds elements to the queue of quantities and basis values of  
// open positions in the pipeline (pending, not closed yet).
func queueAdd(queue, addElem []Pending) []Pending {
	// Switching depending on the lengths of lists of elements 
	// to be added and removed. 
	switch {
	case len(addElem) > 0:
		// Add elements to the queue.
		// Note: Likely, full control over the way the slice is grown 
		// is not needed, so a built-in append function is good enough.
		return append(queue, addElem...)

	default:
		// Do nothing
		return queue
	}
}

// split splits up the queue of Pending objects into two slices: 
// (1) a slice of objects being removed from the queue and 
// (2) a slice of objects remaining in the queue.
func split(queue []Pending, n int) ([]Pending, []Pending) {
	var removed []Pending
	if n > 0 {
		removed = queue[:n]
		queue   = queue[n:]
	}
	return removed, queue
}
