package main

import (
	"fmt"
	"os"
	"strconv"
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
	overallOdds := calculateOdds(odds)
	fmt.Println()
	fmt.Println("Odds:", fmt.Sprintf("%.2f", overallOdds-float64(1))+" to 1", "("+fmt.Sprintf("%.2f", overallOdds)+" for 1"+")")
	fmt.Println("$100 to win:", "$"+strconv.FormatFloat((float64(100)*overallOdds)-float64(100), 'f', 2, 64))
	fmt.Println("Total win:", "$"+strconv.FormatFloat(float64(100)*overallOdds, 'f', 2, 64))
	fmt.Println()
}

func calculateOdds(odds []int) float64 {
	var totalOdds float64 = 1

	if len(odds) == 1 {
		return singleBetOdds(odds[0])
	}

	for _, odd := range odds {
		totalOdds *= singleBetOdds(odd)
	}
	return totalOdds
}

func singleBetOdds(line int) float64 {
	lineFloat := float64(line)
	amountToWager := float64(100)
	var totalWinAmount float64

	// line represents an underdog
	if line > 0 {
		totalWinAmount = amountToWager + lineFloat
		return totalWinAmount / amountToWager
	}

	// line represents a favorite
	totalWinAmount = -lineFloat + amountToWager
	return totalWinAmount / -lineFloat
}
