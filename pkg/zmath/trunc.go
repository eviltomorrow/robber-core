package zmath

import "math"

var n10 = math.Pow10(2)

// Trunc2 trunc2
func Trunc2(val float64) float64 {
	return math.Trunc((val+0.5/n10)*n10) / n10
}
