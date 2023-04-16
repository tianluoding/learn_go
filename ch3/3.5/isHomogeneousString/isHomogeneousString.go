package main

import "fmt"

func isHomogenerousString(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	tmp_a := []rune(a)
	tmp_b := []rune(b)
	mp := make(map[rune]int)
	for _, c := range tmp_a {
		mp[c] = 1
	}
	for _, c := range tmp_b {
		if _, ok := mp[c]; !ok {
			return false
		}
	}
	return true
}

func main() {
	re := isHomogenerousString("田螺姑娘11", "姑娘田螺1")
	fmt.Println(re)
}

