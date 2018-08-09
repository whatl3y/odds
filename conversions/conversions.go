package conversions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RoundingPrecision determines how precise fractional odds will be.
var RoundingPrecision, _ = strconv.Atoi(GetDefaultEnv("ROUNDING_PRECISION", "4"))

// CalculateOdds takes all American odds provided (i.e. -110, 124, etc.)
// and generates the aggregate odds of winning a bet.
func CalculateOdds(odds []int) float64 {
	var totalOdds float64 = 1

	if len(odds) == 1 {
		return SingleBetOdds(odds[0])
	}

	for _, odd := range odds {
		totalOdds *= SingleBetOdds(odd)
	}
	return totalOdds
}

// SingleBetOdds takes an American odds line and converts it to
// the odds that can be used in mathematical operations to determine
// how much can be won given a wager amount.
func SingleBetOdds(americanLine int) float64 {
	lineFloat := float64(americanLine)
	amountToWager := float64(100)
	var totalWinAmount float64

	// line represents an underdog
	if americanLine > 0 {
		totalWinAmount = amountToWager + lineFloat
		return totalWinAmount / amountToWager
	}

	// line represents a favorite
	totalWinAmount = -lineFloat + amountToWager
	return totalWinAmount / -lineFloat
}

// GetAmericanOddsFromOverall converts the overall odds, as returned from
// CalculateOdds and creates the string representation of the corresponding
// American odds.
func GetAmericanOddsFromOverall(overallOdds float64) string {
	overallOdds = overallOdds - 1
	if overallOdds < 1 {
		return "-" + fmt.Sprintf("%.0f", 100/overallOdds)
	}

	return "+" + fmt.Sprintf("%.0f", overallOdds*100)
}

// GetFractionallOddsFromOverallOdds converts the overall odds, as returned from
// CalculateOdds and returns a string of the fractional odds.
func GetFractionallOddsFromOverallOdds(overallOdds float64) string {
	oddsToOne := overallOdds - float64(1)
	oddsToOneRounded := fmt.Sprintf("%."+strconv.Itoa(RoundingPrecision)+"f", oddsToOne)
	fOddsToOneRounded, _ := strconv.ParseFloat(oddsToOneRounded, 64)
	roundingPrecisionMultiple, _ := strconv.Atoi("1" + strings.Join(make([]string, RoundingPrecision+1), "0"))

	oddsToOneInt := int(fOddsToOneRounded * float64(roundingPrecisionMultiple))
	gcd := GCDRemainderRecursive(oddsToOneInt, roundingPrecisionMultiple)
	return strconv.Itoa(oddsToOneInt/gcd) + "/" + strconv.Itoa(roundingPrecisionMultiple/gcd)
}

// GCDRemainderRecursive finds the greatest common denomenator
// given two integers.
func GCDRemainderRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDRemainderRecursive(b, a%b)
}

// CalculateArbitrageProfitMargin takes a slice of American odd integers
// and determines the overall profit margin (if any). Anything negative means
// it is impossible to place an Arbitrage bet, whereas anything positive
// means you can place an Arbitrage bet and be guaranteed to win as long
// as one wager in the set is guaranteed to win.
func CalculateArbitrageProfitMargin(odds []int, extras ...float64) (float64, float64, map[int]float64) {
	wagerAmount, _ := strconv.ParseFloat(GetDefaultEnv("WAGER_AMOUNT", "100.0"), 64)
	if len(extras) > 0 {
		wagerAmount = extras[0]
	}

	individualWagers := make(map[int]float64)
	totalWagered := wagerAmount
	min, _ := minMax(odds)
	minOddsTotalWinAmount := floorToNDecimals(wagerAmount * SingleBetOdds(min))
	alreadySkippedMin := false

	for _, odd := range odds {
		oddsDecimal := SingleBetOdds(odd)
		thisOddWagerAmount := floorToNDecimals(minOddsTotalWinAmount / oddsDecimal)
		individualWagers[odd] = thisOddWagerAmount
		if odd != min || alreadySkippedMin {
			totalWagered += thisOddWagerAmount
		} else {
			alreadySkippedMin = true
		}
	}

	return minOddsTotalWinAmount - totalWagered, totalWagered, individualWagers
}

// GetDefaultEnv gets an environment variable if present, otherwise
// uses a default fallback provided.
func GetDefaultEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func floorToNDecimals(f float64, p ...int) float64 {
	multiplier := 100.0
	if len(p) > 0 {
		intMulitiplier, _ := strconv.Atoi("1" + strings.Join(make([]string, p[0]+1), "0"))
		multiplier = float64(intMulitiplier)
	}

	return float64(int(f*multiplier)) / multiplier
}

func minMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
