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
)

// Trades for signals read from a CSV file; it contains 
// Close (last) price, Trade price (Tx) and the Size of accumulated position 
type Trades struct {
	// Bar ID, such as date/time stamp in any convenient format, 
	// e.g. 'YYYY-MM-DD'
	Dt string

	// Close (last) price
	Cl string

	// Trade price
	Tx string
	
	// The side and the size of position
	SL string
}

// getTrades reads the CSV data file into a slice of Trades objects. 
// No data validation.
func getTrades(file string, headers bool) ([]Trades, error) {
	var (
		sig []Trades
	)
	raw, err := csv2data(file, headers)
	if err != nil {
		msg := "CSV data error!"
		warning(msg, err)
		return sig, err
	}
	sig = data2trades(raw)
	// fmt.Printf("Read signals into the object of type: %T\n", sig)
	return sig, nil
}

// csv2data reads file and puts csv data into a [][]string matrix (raws-columns)
func csv2data(file string, headers bool) ([][]string, error) {
	fmt.Printf("Reading trade signals from file %s ...\n", file)
	csvData, err := readCSV(file)
	if err != nil {
		msg := "Failed to read data from a trade signal file!"
		warning(msg, err)
		return csvData, err
	}
	fmt.Printf("First row in %s: %v\n", file, csvData[0])

	switch headers {
	case true:
		fmt.Printf("Deleting first row (assumed column titles): %v\n", csvData[0])

		// Delete the title row (first element from the slice)
		csvData = append(csvData[:0], csvData[1:]...)
		fmt.Printf("First data row: %v\n", csvData[0])
		return csvData, nil

	default:
		fmt.Printf("First data row: %v\n", csvData[0])
		return csvData, nil
	}
}

// data2trades puts data from a matrix of read input into a slice of 
// Trades objects. No data validation.
func data2trades(csvData [][]string) []Trades {
	var (
		one Trades
		all []Trades
	)
	all = make([]Trades, len(csvData))
	for i, each := range csvData {
		one = Trades{
			Dt: each[0], //col1
			Cl: each[1], //col2
			Tx: each[2], //col3
			SL: each[3], //col4
		}
		all[i] = one
	}
	return all
}

// readCSV reads data from each csv file into a [][]string matrix
func readCSV(filename string) ([][]string, error) {
	var csvData [][]string

	csvFile, errOpen := os.Open(filename)
	if errOpen != nil {
		msgOpen := "Failed to open a trade signal file!"
		warning(msgOpen, errOpen)
		return csvData, errOpen
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, errRead := reader.ReadAll()
	if errRead != nil {
		msgRead := "Failed to read a trade signal file!"
		warning(msgRead, errRead)
		os.Exit(1)
		return csvData, errRead
	}
	return csvData, nil
}
