package safemath

func RandTimeSleep(timeMin float64, timeMax float64) float64 {
	if timeMax == 0 && timeMin == 0 {
		return 0
	}

	if timeMax == timeMin {
		return timeMax
	}

	randTime := RandomFloatRound(timeMin, timeMax, 3)

	return randTime
}

func UpDown(up, num int) []int {
	tmp := make([]int, 0)
	i := 0
	j := 0

	down := num - up

	for k := 0; k < num; k++ {
		upDown := RandomInt(0, 1)

		if upDown == 0 {
			if i < up {
				tmp = append(tmp, 0)
			} else {
				tmp = append(tmp, 1)
			}
			i++
		} else {
			if j < down {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, 0)
			}
			j++
		}
	}

	return tmp
}
