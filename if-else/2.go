// Sana: 17-92-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	fmt.Print("c = ")
	fmt.Scan(&c)

	if a > b && b >= c {
		fmt.Println("a eng kattasi ", a)
	} else if b > a && a >= c {
		fmt.Println("b eng kattasi ", b)
	} else if c > a && a >= b {
		fmt.Println("c eng kattasi ", c)
	} else if a >= b && b > c {
		fmt.Println("a eng kattasi ", a)
	} else if b >= c && c > a {
		fmt.Println("b eng kattasi ", b)
	} else if c >= a && a > b {
		fmt.Println("c eng kattasi ", c)
	} else {
		fmt.Println("Sonlar o'zaro teng: ", a, " ", b, " ", c)
	}
}
