package safemath

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min+1) + min
	}
}

func RandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	if min > max {
		return min
	} else {
		return rand.Float64()*(max-min) + min
	}
}

func RandomFloatRound(min, max, prec float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())

	answer := 0.0

	if min > max {
		answer = Round(min, prec)
	} else {
		answer = Round(rand.Float64()*(max-min)+min, prec)
	}

	return answer
}
