package zmath

func MinFloat64(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var min = data[0]
	for i := 1; i <= len(data)-1; i++ {
		if data[i] < min {
			min = data[i]
		}
	}
	return min
}
