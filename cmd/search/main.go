package main

import "fmt"

type item struct {
	name  string
	value int
	// hoge fuga
	// ...
}

func main() {
	items := []item{{"a", 167}, {"b", 3567}, {"c", 14678}, {"d", 4561}, {"e", 1345}, {"f", 879}, {"g", 5346}, {"h", 8796}, {"i", 21}, {"j", 45}, {"k", 87}}
	fmt.Println(items)

	fmt.Println(forSearch(items, "k"))
	fmt.Println(forSearch(items, "x"))

	m := make(map[string]int, len(items))
	for _, v := range items {
		m[v.name] = v.value
	}
	fmt.Println(m)

	fmt.Println(mapSearch(m, "k"))
	fmt.Println(mapSearch(m, "x"))
}

func forSearch(items []item, target string) int {
	for _, v := range items {
		if v.name == target {
			return v.value
		}
	}
	return -1000
}

func mapSearch(m map[string]int, target string) int {
	if v, flag := m[target]; flag {
		return v
	}
	return -1000
}
