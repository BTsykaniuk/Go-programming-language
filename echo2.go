package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	s, sep := "", ""

	fmt.Println(os.Args[1:])
	for i, arg := range os.Args[1:]{
		s += sep + arg + sep
		sep = " "
		fmt.Println(i)
	}

	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
