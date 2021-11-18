package zmath

import (
	"fmt"
	"testing"
)

func TestLeastSquares(t *testing.T) {
	var y = []float64{50.3, 51.4, 50.9, 49.9, 50.5}
	var x = []float64{1, 1, 1, 1, 1}

	a, b := LeastSquares(x, y)
	fmt.Println(a, b)

	var m = a*6 + b
	fmt.Println(m)
}
