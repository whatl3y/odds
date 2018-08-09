package main

import (
	"fmt"
	"os"
	"strconv"

	c "github.com/whatl3y/odds/conversions"
)

func main() {
	var odds []int
	for _, sOdd := range os.Args[1:] {
		odd, err := strconv.Atoi(sOdd)
		if err != nil {
			panic(err)
		}
		odds = append(odds, odd)
	}
	expectedReturn, totalWagered, individualWagers := c.CalculateArbitrageProfitMargin(odds)

	// Show the profit margin and expected return from a 100 base unit
	// Anything positive for expected return provides an opportunity
	// for betting arbitrage, otherwise there is a net negative return
	fmt.Printf(`
%s base unit
Individual Wagers: %v
Total Wagered: %f
Expected return: %f
Profit Margin: %f%%

`, c.GetDefaultEnv("WAGER_AMOUNT", "100"), individualWagers, totalWagered, expectedReturn, (expectedReturn/totalWagered)*100)
}
