package main

import "fmt"

func main() {
	a := [...]int{99: -1}
	fmt.Println(len(a))
	var b = [...]int{99:1}
	fmt.Println(a == b)
}
