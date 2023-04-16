package main

import (
	"fmt"
	"github/tianluoding/learn_go/pkg/util"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var fileResult []string
	if len(files) == 0 {
		util.CountLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			f1, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			if util.IsRepeat(f) {
				fileResult = append(fileResult, arg)
			}
			util.CountLines(f1, counts)
			f.Close()
			f1.Close()
		}
	}
	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, lines)
		}
	}
	for _, file := range fileResult {
		fmt.Println(file)
	}

}
