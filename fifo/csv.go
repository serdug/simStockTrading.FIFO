// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	attributes string = `Bar
ClosePx
TradePx
SHORT
LONG
Entry.S
Exit.S
Entry.L
Exit.L
Quantity.S
Quantity.L
Basis.S
Basis.L
Realized.S
Realized.L
Assets`
)

// writeCSVbasic exports results of calculations in the CSV format
func writeCSVbasic(allRecords []Asset, outFile string) {
	var (
		csvNewTName string = outFile
		err1 error
		field []string
	)

	csvNewFile, err := os.OpenFile(csvNewTName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Creating output file:", csvNewTName)
		csvNewFile, err1 = os.Create(csvNewTName)
		if err1 != nil {
			fmt.Println("Output file creating error:", err1)
			return
		}
	}
	defer csvNewFile.Close()

	writer := csv.NewWriter(csvNewFile)

	headers := strings.Split(attributes, "\n")
	writer.Write(headers)
	// fmt.Println("Headers:", len(headers))

	for _, one := range allRecords {
		field = make([]string, len(headers))

		field[0] = one.Bar
		// Convert response to a single string.
		field[1] = fmt.Sprintf("%f", one.Pxs.Cl)
		field[2] = fmt.Sprintf("%f", one.Pxs.Tx)
		field[3] = strconv.Itoa(-one.S.Pos.E)
		field[4] = strconv.Itoa(+one.L.Pos.E)
		field[5] = strconv.Itoa(-one.S.Pos.I)
		field[6] = strconv.Itoa(+one.S.Pos.O)
		field[7] = strconv.Itoa(+one.L.Pos.I)
		field[8] = strconv.Itoa(-one.L.Pos.O)
		field[9] = fmt.Sprintf("%f", one.S.Qty.E)
		field[10] = fmt.Sprintf("%f", one.L.Qty.E)
		field[11] = fmt.Sprintf("%f", one.S.Basis.E)
		field[12] = fmt.Sprintf("%f", one.L.Basis.E)
		field[13] = fmt.Sprintf("%f", one.S.Result.Rzd)
		field[14] = fmt.Sprintf("%f", one.L.Result.Rzd)
		field[15] = fmt.Sprintf("%f", one.NAV)

		writer.Write(field)
	}
	writer.Flush()
	return
}
