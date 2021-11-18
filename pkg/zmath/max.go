package zmath

func MaxFloat64(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	var max float64
	for _, d := range data {
		if d > max {
			max = d
		}
	}
	return max
}
