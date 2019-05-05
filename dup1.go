package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	countWords := make(map[string]int)
	inputText := bufio.NewScanner(os.Stdin)

	for inputText.Scan() {
		countWords[inputText.Text()] ++
	}
	fmt.Println(countWords)

	for k, v := range countWords {
		if v > 1 {
			fmt.Println(k)
		}
	}
}

