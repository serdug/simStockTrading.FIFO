# Simulated Stock Trading – First-In-First-Out

A fast stock trading performance calculator for backtesting and simulating stock trading strategies in Go. 

The calculation reveals Net Asset Value dynamics, given that trade signals – side and size – and prices are provided. It works locally with configuration and parameters set through a YAML file; tested on Unix-like (Linux and MacOS) operating systems.


## Accounting Policy

* No reinvestments of profits
* Constant limit of exposure per position applies; no position is rebalanced between its entry and its exit
* Positions are opened and closed following the First-In-First-Out (FIFO) order
* The calculation principles generally resemble the rules which Interactive Brokers use for customer account reporting
* The 'short' and the 'long' sides are being treated independently
* Quantities are not rounded to integer values, as this enables calculations for securities undergoing splits (and reverse splits)


## Settings

Settings are read from a YAML file containing the configuration of input data and calculation parameters. A full name of the config file `a/path/to/file.yaml` must be passed as the mandatory argument.


## Input

Stock prices and trade signals, such as position's side and size, are passed to the calculator in the CSV format. The respective CSV files should be listed as `signals` in the config file. 

**The input data is taken 'as is' with no validation whatsoever.**

There's no need to name columns exactly as shown below (or name the columns at all if `headers: no`). However, the input data should be arranged in 4 columns and the following column order must be respected:

* `Bar` (character string) – bar ID, e.g. date/time stamp, in any convenient form; any values, unique and duplicate, are allowed
* `Close Price` (numeric) – stock's Close price
* `Trade Price` (numeric) – stock's Trade price, a price at which opening or closing trades would be executed; this is an assumed gross price, before broker's fees would be added/deducted
* `Position` (integer) - the required position size, negative for 'short', positive for 'long' positions


## Titles

If the settings indicate that the first row contains column titles (`headers: yes`), the first row of every input file is ignored. 


## Output

The results of calculation are returned in CSV files. Locations and names of output files should be listed under `results` in the config file. The length of this list must be the same as the length of the `signals` list.

The basic output format:

* `Bar` (character string) – bar ID as per input data
* `ClosePx` (numeric) – stock's Close price as per input data
* `TradePx` (numeric) – stock's Trade price as per input data
* `SHORT` (integer, 0 or negative) – the accumulated position of all 'short' entries and exits
* `LONG` (integer, 0 or positive) – the accumulated position of all 'long' entries and exits
* `Entry.S` (integer, 0 or negative) - a signal to sell stocks opening a 'short' position of the indicated size
* `Exit.S` (integer, 0 or positive) - a signal to buy stocks closing a 'short' position of the indicated size
* `Entry.L` (integer, 0 or positive) - a signal to buy stocks opening a 'long' position of the indicated size
* `Exit.L` (integer, 0 or negative) - a signal to sell stocks closing a 'long' position of the indicated size
* `Quantity.S` (numeric, negative) - the resulting 'short' stock quantity, obtained by adding up the product of the limit of exposure per position and the size of position divided by its net cost price, for each existing 'short' position
* `Quantity.L` (numeric, positive) - the resulting 'long' stock quantity (the calculation is similar to that of the 'short' stock quantity)
* `Basis.S` (numeric) - the resulting 'short' stock basis, obtained by adding up the product of the 'short' stock quantity and its net cost price, for each existing 'short' position
* `Basis.L` (numeric) - the resulting 'long' stock basis (the calculation is similar to that of the 'short' stock basis)
* `Realized.S` (numeric) - the 'short' trade return realized at a time
* `Realized.L` (numeric) - the 'long' trade return realized at a time
* `Assets` (numeric) - the resulting Net Asset Value, includes realized and unrealized returns


## Parameters

Same parameters for all inputs:

* `cash` (numeric) - cash initially allocated for trading
* `limit` (numeric) - the limit of exposure (USD) per position
* `commission` (numeric) - broker's commission payable per share bought or sold


## Dependencies

* `go`
* `fmt`
* `flag`
* `os`
* `io/ioutil`
* `encoding/csv`
* `gopkg.in/yaml.v2`
* `strconv`
* `strings`
* `math`
* `time`


***



## Paths

Absolute paths to files are required in `config.yaml`, for instance:

```{yaml}
input: "/home/<USER_NAME>/io.calc/out/results.csv"
```


## Compilation and Running

Pre-compile and run

* Pre-compile.

```
go install github.com/serdug/simStockTrading.FIFO
```

* Create folders for settings, signals and output files. It may be convenient to use local folders `~/io.calc/settings/`, `~/io.calc/in/` and `~/io.calc/out/` for settings, signals and output results, respectively.

* Copy example files in the respective folders.

* Update the `config-fifo.yaml` file accordingly.

* Run a compiled package. A full name, including path, of the config file should be passed as an argument.

```
$GOPATH/bin/simStockTrading.FIFO ~/io.calc/settings/config-fifo.yaml
```


## License

[MIT](https://github.com/serdug/simStockTrading.FIFO/blob/master/LICENSE)
