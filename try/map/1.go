/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Slicening har bir elementi nechtadadan takrorlanganligini topuvchi dastur tuzing
[1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 3, 4, 5, 6, 6]
[1:6 2:5 3:1 4:1 5:1 6:2]
*/

package main

import "fmt"

func main() {
	var (
		n int
	)

	fmt.Print("Slice uzunligini kiriting: ")
	fmt.Scan(&n)

	org := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&org[i])
	}

	try(org)

}

func try(copy []int) {
	mapp := make(map[int]int)

	for _, v := range copy {
		mapp[v]++
	}
	fmt.Println(mapp)
}
