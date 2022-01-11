package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 3, 5, 9}
	comb := combination(s)
	fmt.Println(comb)
}

func combination(s []int) [][2]int {
	l := len(s)
	comb := make([][2]int, 0, l*(l-1)/2)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			comb = append(comb, [2]int{s[i], s[j]})
		}
	}
	return comb
}
