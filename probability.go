package safemath

func PercentProbability(percent int) int {
	if RandomInt(0, 100) < percent {
		return 1
	}

	return -1
}

func PercentProbabilityFloat(percent float64) float64 {
	if RandomFloat(0, 100) < percent {
		return 1
	}

	return -1
}
