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
	overallOdds := c.CalculateOdds(odds)

	// Print out all the different representations/formats of odds
	// and provide an example of what should be won given a $100 wager.
	fmt.Printf(`
American Odds: %s
Decimal Odds: %s
Fractional Odds: %s
General Odds: %s to 1 (%s for 1)

$100 to win: $%s
Total win: $%s

`,
		c.GetAmericanOddsFromOverall(overallOdds),
		fmt.Sprintf("%.2f", overallOdds),
		c.GetFractionallOddsFromOverallOdds(overallOdds),
		fmt.Sprintf("%.2f", overallOdds-float64(1)),
		fmt.Sprintf("%.2f", overallOdds),
		strconv.FormatFloat((float64(100)*overallOdds)-float64(100), 'f', 2, 64),
		strconv.FormatFloat(float64(100)*overallOdds, 'f', 2, 64),
	)
}
