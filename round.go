package safemath

import "math"

func Round(x float64, prec float64) float64 {
	var rounder float64
	pow := math.Pow(10, prec)
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	if prec < 0 {
		p := math.Pow(10, -1*prec)
		return rounder * p
	} else {
		return rounder / pow
	}
}
