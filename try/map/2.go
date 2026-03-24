/*
Sana: 21-03-2026
Muallif: Izzatillayev Komronbek
Maqsad: Slicening eng kop takrorlangan elementini toping
[1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 3, 4, 5, 6, 6]
[1:6 2:5 3:1 4:1 5:1 6:2]
*/

package main

import "fmt"

func main() {
	var arr []int = []int{1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 3, 4, 5, 6, 6}

	try1(arr)
}

func try1(copy []int) {
	mapp := make(map[int]int)

	for _, v := range copy {
		mapp[v]++
	}

	katta := mapp[1]

	for i := 1; i <= 6; i++ {
		for j := 2; j <= 6; j++ {
			if mapp[j] > mapp[i] {
				katta = j
			} else {
				katta = i
			}
		}
	}

	fmt.Println(katta)
}
