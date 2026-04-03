package myMath

import "fmt"

func Trunc(x float64) int {
	return int(x) / 1
}
func Pi() float64 {
	p := 31.4159 / (5 * 2)
	return p
}
func Log(a, b int) int {
	if a == 1 || b < 0 || a < 0 {
		fmt.Println("xato son kiritildi!")
		return 0
	}
	mult := 1
	count := 0
	for {

		r := mult * a
		mult = r
		count++
		if r == b {
			return count
		}
	}
}
func Log10(x int) int {
	if x < 0 {
		fmt.Println("xato son kiritildi!")
		return 0
	}
	mult := 1
	count := 0
	for {

		r := mult * 10
		mult = r
		count++
		if r == x {
			return count
		}
	}
}
func Pow10(x int) int {
	a := 10
	mult := 1
	r := 0
	for a > 0 {
		r = mult * x
		mult = r
		a--
	}
	return r
}
