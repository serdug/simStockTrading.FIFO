// Copyright (c) 2020 Sergey Dugaev. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file in the project root for more information.

// Package fifo models the First-In-First-Out position management 
// to calculate results of algorithmic trading by trade signals, 
// given that returns are not reinvested and positions are not rebalanced.
package fifo

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration values
type Config struct {
	// The location of home folder
	Home    string   `yaml:"home"`

	// Input file names
	Signals []string `yaml:"signals"`

	// A flag showing whether input files contain column titles in the first rows
	Headers bool     `yaml:"headers"`

	// Output file names
	Results []string `yaml:"results"`

	// Starting asset value, cash initially allocated for trading
	Cash    float64  `yaml:"cash"`

	// Cash limit of exposure per position
	Lim     float64  `yaml:"limit"`

	// Broker commission
	Fee     float64  `yaml:"commission"`
}

// Params is the object for parameters
type Params struct {
	// Starting asset value, cash initially allocated for trading
	Cash    float64

	// Cash limit of exposure per position
	Lim     float64

	// Broker's commission
	Fee     float64
	
	// A flag showing whether input files contain column titles in the first row
	Headers bool
}

// ReadConfig parses a YAML config file .
// File name as the first argument is expected.
func ReadConfig() Config {
	var conf Config

	configFile := argument(0)
	if len(configFile) == 0 {
		fmt.Println("\nNo config file found! A full name ('a/path/to/file.yaml') of the config file must be passed as the argument.")
		return conf
	}
	fmt.Println("\nReading config file:", configFile)

	dat, errReadFile := ioutil.ReadFile(configFile)
	msgReadFile := "Config read failed!" 
	warning(msgReadFile, errReadFile)

	errYaml := yaml.Unmarshal(dat, &conf)
	msgYaml := "YAML unmarshal failed!" 
	warning(msgYaml, errYaml)
	return conf
}

// argument reads and returns the required argument. 
func argument(n int) string {
	flag.Parse()
	return flag.Arg(n)
}
