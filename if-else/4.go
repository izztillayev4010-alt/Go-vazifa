// Sana: 17-92-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!

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
		fmt.Println("eng kichigi ", c)
	} else if b > a && a >= c {
		fmt.Println("eng kichigi ", c)
	} else if c > a && a >= b {
		fmt.Println("eng kichigi ", b)
	} else if a >= b && b > c {
		fmt.Println("eng kichigi ", c)
	} else if c >= b && b > a {
		fmt.Println("eng kichigi ", a)
	} else if b > c && c > a {
		fmt.Println("eng kichigi ", a)
	} else if c >= a && a > b {
		fmt.Println("c eng kichigi ", b)
	} else {
		fmt.Println("Sonlar o'zaro teng: ", a, " ", b, " ", c)
	}
}
