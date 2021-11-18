package zmath

import (
	"fmt"
	"testing"
)

func TestDiscountForPV(t *testing.T) {
	var n = 10
	var i = 0.149933
	var pmt = 7.5
	var fv = 9.5
	var pv = PV(n, i, pmt, fv)
	fmt.Printf("%v\r\n", pv)
}

func TestCalDiscountForFV(t *testing.T) {
	var n = 4
	var i = 0.15
	var pmt = 2000000.0
	var pv = 5710000.0
	var fv = FV(n, i, pmt, pv)
	fmt.Println(fv)
}

func TestIRR(t *testing.T) {
	var n = 10
	var pmt = 2.5
	var pv = 40.0
	var fv = 9.5
	var i = IRR(n, pmt, pv, fv)
	fmt.Printf("%v\r\n", i)
}

func TestJDIRR(t *testing.T) {
	var n = 12
	var pmt = 537.4
	var pv = 5999.0
	var fv = 0.0
	var i = IRR(n, pmt, pv, fv)
	fmt.Printf("%v\r\n", i)
}
