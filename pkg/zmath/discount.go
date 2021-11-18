package zmath

func PV(n int, i float64, pmt float64, fv float64) float64 {
	for ; n > 0; n-- {
		fv += pmt
		fv = fv / (1 + i)
	}
	return Trunc2(fv)
}

func FV(n int, i float64, pmt float64, pv float64) float64 {
	for m := 0; m < n; m++ {
		pv = pv*(1+i) - pmt
	}
	return Trunc2(pv)
}

func IRR(n int, pmt float64, pv float64, fv float64) float64 {
	var (
		precision = 0.000001
		loopLimit = 1000
		arr2      = 2 * (pmt*float64(n) + fv - pv) / pv / float64(n)
		min, max  = 0.0, arr2
		guess     = 0.0
	)

	for {
		if loopLimit < 0 {
			return -1
		}
		loopLimit--

		guess = (min + max) / 2
		var val = pv
		for i := 0; i < n; i++ {
			val = val*(1+guess) - pmt
		}
		val = val - fv

		if val > 0 && val < precision {
			break
		}
		if val < 0 {
			min = guess
		}
		if val > 0 {
			max = guess
		}
	}
	return Trunc2(guess * 100)
}
