package safemath

import (
	"math/rand"
	"time"
)

func volumeRandomizer(volume, minAmount, maxAmount float64, length int, result *[]float64) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var a []float64
	aSum := 0.
	for i := 0; i < length; i++ {
		v := r.Float64()
		a = append(a, v)
		aSum += v
	}

	delta := maxAmount - minAmount

	aAvgMin := (volume - minAmount*float64(length)) / (delta * float64(length))

	if aSum < aAvgMin*float64(length) {
		aSum = 0.
		for i := range a {
			a[i] = a[i] - a[i]*aAvgMin + aAvgMin
			aSum += a[i]
		}
	}

	delta_ := (volume - minAmount*float64(length)) / aSum

	for i := 0; i < length; i++ {
		*result = append(*result, minAmount+delta_*a[i])
	}
}

func RandAmountsForPeriod(minAmount, amount float64, minTrades, maxTrades int) (int, []float64) {
	numTrade := RandomInt(minTrades, maxTrades)

	var amountsForTrade []float64
	if minAmount > 0.9*amount/float64(numTrade) {
		volumeRandomizer(amount, 0.0, amount, numTrade, &amountsForTrade)
	} else {
		volumeRandomizer(amount, minAmount, amount, numTrade, &amountsForTrade)
	}

	return numTrade, amountsForTrade
}
