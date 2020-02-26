// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

import (
	"fmt"
)

// Model runs trade result calculations based on the history of trade signals
// passed as data files.
func Model(sigFile, outFile string, par Params) {
	fmt.Printf("\nHeaders (%T): %v\n", par.Headers, par.Headers)
	
	sigs, errSig := getTrades(sigFile, par.Headers)
	if errSig != nil {
		msgSig := "Signal read failed!"
		warning(msgSig, errSig)
		return
	}

	// Parameters:
	fmt.Printf("Cash initially allocated for trading (%T): %v\n", par.Cash, par.Cash)
	fmt.Printf("Limit of exposure per position (%T)      : %v\n", par.Lim, par.Lim)
	fmt.Printf("Fee (%T): %v\n", par.Fee, par.Fee)

	results, errFIFO := fifo(argsFIFO{
		Sigs:     sigs,
		Cashbase: par.Cash,
		Lim:      par.Lim,
		Fee:      par.Fee,
	})
	if errFIFO != nil {
		msgFIFO := "Ups-a-daisy... Calculation failed!"
		warning(msgFIFO, errFIFO)
		return
	}

	// fmt.Println("Records in results:", len(results))
	fmt.Println("Starting NAV:", results[0].NAV)
	fmt.Printf("Ending NAV  : %v on %s\n", results[len(results)-1].NAV, results[len(results)-1].Bar)
	fmt.Printf("Number of finished trades: %v on %s\n", results[len(results)-1].ExitN, results[len(results)-1].Bar)

	writeCSVbasic(results, outFile)
}

// warning prints an error message. It does not cause the process to end.
func warning(msg string, e error) {
	if e != nil {
		fmt.Printf("%s [%v]\n", msg, e)
	}
}
