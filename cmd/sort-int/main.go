package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 3, 8, 2, 6}
	fmt.Println(s)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println(s)
	for _, v := range []int{7, 3, 9, -1} {
		s = insert(s, v)
		fmt.Println(s)
	}
}

func insert(s []int, e int) []int {
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	if i == len(s) {
		return append(s, e)
	}
	s = append(s[:i+1], s[i:]...)
	s[i] = e
	return s
}
