package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src := os.Args[1]
	if src == "" {
		fmt.Println(0)
	} else {
		count := len(strings.Split(src, " "))
		fmt.Println(count)
	}
}
