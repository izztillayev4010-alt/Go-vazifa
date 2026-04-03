package myMath

func Pow(x, n int) int {
	mult := 1
	var r int
	if x == 0 {
		return 0
	} else {
		for i := 1; i <= n; i++ {

			r = mult * x
			mult = r
		}
	}
	return r

}

func Sqrt(x int) float64 {
	org := x
	for i := 0; i < 10; i++ {
		if x > 0 {
			x = (x + org/x) / 2
		} else {
			return -1.0
		}
	}
	return float64(x)

}
func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
