package zmath

func SumInt64(data []int64) int64 {
	if len(data) == 0 {
		return 0
	}

	var sum int64
	for _, d := range data {
		sum += d
	}
	return sum
}

func SumFloat64(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum
}
