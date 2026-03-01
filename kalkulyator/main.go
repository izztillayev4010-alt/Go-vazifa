package main

import (
	"errors"
	"fmt"
)

func main() {
	var r, res float64

	fmt.Println("Quyidagi amallarda birini tanlang: 1)arifmetik amallar, 2)foizga doir amallar")
	fmt.Scan(&r)

	switch r {
	case 1:
		res = arifmetik()
		fmt.Println(res)
	case 2:
		res1 := foiz()
		fmt.Println(res1)
	}

}
func arifmetik() float64 {
	fmt.Println("Quyidagi amallardan birini tanlang: +, -, *, / ")

	var q string
	var res float64 = 0

	for {
		fmt.Scan(&q)
		if q == "=" {

			break
		}

		var a, b float64
		fmt.Print("a = ")
		fmt.Scan(&a)
		fmt.Print("b = ")
		fmt.Scan(&b)

		switch q {
		case "+":
			res += add(a, b)
		case "-":
			res += sub(a, b)
		case "*":
			res = multi(a, b)
		case "/":
			res1, err := div(a, b)
			if err != nil {
				fmt.Println(err)
			}
			res = res1

		}

	}
	return res
}
func add(x float64, y float64) float64 {
	return x + y
}
func sub(x float64, y float64) float64 {
	return x - y
}
func multi(x float64, y float64) float64 {
	return x * y
}
func div(x float64, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("Nolga bo'lib b'lmaydi!")
		return 0, err
	}
	return x / y, nil
}
func foiz() int {
	fmt.Println("Quyidagilardan birini tanlang: 1) Sonning p foizini topish, 2)Sonni p foizga oshirish 3)Sonni p foizga kamaytirish,4)Sonni x ga bolgandagi qoldiqni topish")
	var w int
	fmt.Scan(&w)

	var s, p, res int
	fmt.Print("Sonni kiriting: ")
	fmt.Scan(&s)

	fmt.Print("Foizini kiriting: ")
	fmt.Scan(&p)

	switch w {
	case 1:
		res = ftop(s, p)
	case 2:
		res = fosh(s, p)
	case 3:
		res = fkam(s, p)
	case 4:
		res = qol(s, p)
	}
	return res
}
func ftop(x int, y int) int {
	return x * y / 100
}
func fosh(x int, y int) int {
	return x + x*y/100
}
func fkam(x int, y int) int {
	return x + x*y/100
}
func qol(x int, y int) int {
	return x % y
}
