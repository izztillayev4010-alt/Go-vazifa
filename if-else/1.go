// Sana: 17-92-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: a va b sonlari berilgan. Bu sonlarning kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("a = ")
	fmt.Scan(&a)

	fmt.Print("b = ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Kattasi a ", a)
	} else if b > a {
		fmt.Println("Kattasi b ", b)
	} else {
		fmt.Println("Ikkalisi teng")
	}

}
