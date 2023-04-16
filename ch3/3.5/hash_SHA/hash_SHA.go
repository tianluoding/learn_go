package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	input := args[0]
	//fmt.Println(args)
	if len(args) == 1 {
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
	}else {
		method := args[1]
		if method == "SHA256" {
			fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
		}else if method == "SHA384" {
			fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
		}else if method == "SHA512" {
			fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
		} else {
			fmt.Println("input error!")
		}
	}
}
