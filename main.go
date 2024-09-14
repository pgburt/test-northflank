package main

import (
	"math/big"
	"time"

	"time"
	"fmt"
)

// The Leibniz caculation for Pi/4 follows.
// (pi/4) = 1 - (1/3) + (1/5) - (1/7) ...
// We need to find each new fraction, and its sign.
// Add each new number in a series.

func calcDenominator(i int) int64 {
	// Calculate the denominator.
	// The value should be the nth odd number
	// For ex: if n = 4, then return 7.
	return int64(i*2 - 1)
}

func findSign(i int, d int64) *big.Rat {
	if i%2 == 0 {
		return big.NewRat(-1, d)
	} else {
		return big.NewRat(1, d)
	}
}

func sleepyTime() {
	// Used only to demo call times in Continuous Profiling.
	time.Sleep(10 * time.Millisecond)
}

func main() {
	// Rat is for Rational Numbers of arbitrary length.
	sum := new(big.Rat)
	r := new(big.Rat)
	// The Leibniz caculation for Pi/4 follows.
	// (pi/4) = 1 - (1/3) + (1/5) - (1/7) ...
	for i := 1; i <= 10000; i++ {
		d := calcDenominator(i)
		r = findSign(i, d)
		sum = sum.Add(sum, r)
		sleepyTime()
	}
	// Change pi/4 to just pi by multiplying by 4.
	multiplier := big.NewRat(4, 1)
	sum = sum.Mul(multiplier, sum)

	// FloatString asks for an amount of precision.
	fmt.Println(sum.FloatString(32))
}
