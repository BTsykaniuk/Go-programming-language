package main


import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	wordCount := make(map[string]int)
	lineSource := make(map[string]map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("File not found")
		countLines(os.Stdin, wordCount)
		lineSource["stdin"] = wordCount
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}

			countLines(f, wordCount)
			lineSource[f.Name()] = wordCount

			f.Close()
		}
	}

	for source, countWords := range lineSource {
		for line, count := range countWords {
			if count >= 2 {
				fmt.Printf("Find %d duplicates for line %s in file %s\n", count, line, source)
			}
		}
	}
}


func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()] ++
	}
}