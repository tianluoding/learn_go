package main

import (
	"fmt"
	"github/tianluoding/learn_go/pkg/util"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		util.CountLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			util.CountLines(f, counts)
			f.Close()
		}
	}
	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, lines)
		}
	}

}
