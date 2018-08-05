# odds

Calculate the true odds of a single bet or parlay given a
list of single events' odds in American format (i.e -110, 100, etc.).

## Install

```sh
$ go install github.com/whatl3y/odds
```

## Examples

```sh
# if '$GOPATH/bin' is in your path:
$ odds 100
# Total Odds: 2.00 to 1
# $100 to win: $100.00
# Total win: $200.00

# Multiple lines to calculate parlay odds
$ odds -110 120 -240
# Total Odds: 5.95 to 1
# $100 to win: $495.00
# Total win: $595.00

# if '$GOPATH/bin' is NOT in your path:
$ $GOPATH/bin/odds -110 120 -240 customer_id
# Total Odds: 5.95 to 1
# $100 to win: $495.00
# Total win: $595.00
```