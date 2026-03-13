// Sana: 13-03-2026
// Muallif: Izzatillayev Komronbek
// Maqsad: Array yoki slice elementlari uchida polindromlarini alohida slicega yig’uvchi dastur tuzing!
//(faqat array bilan)

package main

import "fmt"

func main() {

	var number [10]int = [10]int{}
	polindrom := []int{}

	for i := 0; i < len(number); i++ {
		fmt.Scan(&number[i])
	}
	for _, a := range number {

		if ispolindrome(a) {
			polindrom = append(polindrom, a)
		}
	}

	fmt.Println(polindrom)
}

func ispolindrome(number int) bool {

	original := number
	reverse := 0

	for number > 0 {

		remain := number % 10
		reverse = reverse*10 + remain
		number = number / 10
	}

	return original == reverse
}
