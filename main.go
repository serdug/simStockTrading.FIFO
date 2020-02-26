// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

package main

import (
	"fmt"
	"time"

	"github.com/serdug/calc-tradesim-fifo/fifo"
)

var (
	config = fifo.ReadConfig()
)

func main() {
	// *** START THE TIMER ***
	t0 := time.Now().UTC()

	switch {
	case len(config.Signals) == 0:
		// Do nothing
		fmt.Println("No trade signals found! Calculation aborted.\n")

	case len(config.Signals) == len(config.Results):
		// Note: the same parameters for all inputs.
		params := fifo.Params{
			Cash:    config.Cash,
			Lim:     config.Lim,
			Fee:     config.Fee,
			Headers: config.Headers,
		}

		for i := range config.Signals {
			fifo.Model(config.Home + config.Signals[i], 
				config.Home + config.Results[i], 
				params)
		}

	default:
		// Do nothing
		fmt.Println("Check the config! The numbers of input and output files must be the same.")
	}

	// *** STOP THE TIMER ***
	// Integer without decimals
	t := float64(time.Since(t0) / time.Millisecond)

	fmt.Printf("Latency (millisec): %v\n", t)
}
