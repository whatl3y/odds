# odds

This repo contains a couple of CLI tools you can install and run
in addition to several packages that are useful in your project(s)
related to sports wagering odds calculations.

## CLI Tools

### odds

Calculate the true odds of a single bet or parlay given a
list of single events' odds in American format (i.e -110, 100, etc.)

#### Install

```sh
$ go get github.com/whatl3y/odds/odds
```

#### Examples

```sh
# if '$GOPATH/bin' is in your path:
$ odds 100
# Total Odds: 1.00 to 1 (2.00 for 1)
# $100 to win: $100.00
# Total win: $200.00

# Multiple lines to calculate parlay odds
$ odds -110 120 -240
# Total Odds: 4.95 to 1 (5.95 for 1)
# $100 to win: $495.00
# Total win: $595.00

# if '$GOPATH/bin' is NOT in your path:
$ $GOPATH/bin/odds -110 120 -240
# Total Odds: 4.95 to 1 (5.95 for 1)
# $100 to win: $495.00
# Total win: $595.00
```

### arb (Arbitrage Wagering)

Calculate the ability to place an [Arbitrage bet](https://en.wikipedia.org/wiki/Arbitrage_betting)
by placing a wager on all sides of the provided odds as long as one wager in the
provided set of odds guaranteed to win.

If the expected return is positive, it means it is possible to place wagers
between all sides and guarantee a

#### Install

```sh
$ go get github.com/whatl3y/odds/arb
```

#### Examples

```sh
$ arb -110 -110
# 100 base unit
# Total Wagered: 199.990000
# Expected return: -9.090000
# Profit Margin: -4.545227%

$ arb -238 293
# 100 base unit
# Total Wagered: 136.130000
# Expected return: 5.880000
# Profit Margin: 4.319401%
```

## Packages

### conversions

The `conversions` package provides helper methods that are used in the CLI
tools to convert between different odds types, calculate overall odds of
single and multiple/parlay bets, etc.

#### Install

```sh
$ go get github.com/whatl3y/odds/conversions
```

#### API

TODO
