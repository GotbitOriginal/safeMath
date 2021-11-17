package safemath

func PercentProbability(percent int) int {
	if RandomInt(0, 100) < percent {
		return 1
	}

	return -1
}
