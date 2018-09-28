package conversions

import (
	"fmt"
	"testing"
)

func TestCalculateOdds(t *testing.T) {
	overall1 := CalculateOdds([]int{100})
	if overall1 != 2 {
		t.Error("Overall with one entered odd not calculated correctly: expected 1", overall1)
	}

	overall11 := CalculateOdds([]int{-110})
	if fmt.Sprintf("%.2f", overall11) != "1.91" {
		t.Error("Overall with one entered favorite odd not calculated correctly: expected 1.91", overall11)
	}

	overall2 := CalculateOdds([]int{100, 100})
	if overall2 != 4 {
		t.Error("Overall with two entered odd not calculated correctly: expected 4", overall2)
	}
}

func TestGetAmericanOddsFromOverall(t *testing.T) {
	am1 := GetAmericanOddsFromOverall(2)
	if am1 != "+100" {
		t.Error("American odds not returned correctly: expected +100", am1)
	}

	am2 := GetAmericanOddsFromOverall(1.5)
	if am2 != "-200" {
		t.Error("American odds not returned correctly: expected -200", am2)
	}
}

func TestGetFractionalOddsFromOverallOdds(t *testing.T) {
	fr1 := GetFractionalOddsFromOverallOdds(2)
	if fr1 != "1/1" {
		t.Error("Fractional odds not returned correctly: expected 1/1", fr1)
	}

	fr2 := GetFractionalOddsFromOverallOdds(1.5)
	if fr2 != "1/2" {
		t.Error("Fractional odds not returned correctly: expected 1/2", fr2)
	}
}

func TestGCDRemainderRecursive(t *testing.T) {
	gcd := GCDRemainderRecursive(12, 4)
	if gcd != 4 {
		t.Error("GCD not calculated correctly: expected 4", gcd)
	}
}

func TestCalculateArbitrageProfitMargin(t *testing.T) {
	ret, _, _ := CalculateArbitrageProfitMargin([]int{-238, 293})
	if floorToNDecimals(ret, 1) != 5.8 {
		t.Error("Error calculating Arbitrage return on $100 wager: expected 5.8", floorToNDecimals(ret, 1))
	}

	ret9, _, _ := CalculateArbitrageProfitMargin([]int{-110, -111}, 100.0)
	if ret9 > -9.0 {
		t.Error("Error calculating Arbitrage return on $100 wager: expected return < -9", ret9)
	}
}
