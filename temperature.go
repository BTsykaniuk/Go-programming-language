package main


import (
	"fmt"
	"learn/tempconverter"
)


func main() {
	fmt.Printf("Very cool - %v\n", tempconverter.AbsoluteNullC)
	fmt.Printf("Very cool in K - %v\n", tempconverter.CtoK(tempconverter.AbsoluteNullC))
	fmt.Printf("Hot - %v\n", tempconverter.BoolingC)
	fmt.Printf("Hot in K - %v\n", tempconverter.CtoK(tempconverter.BoolingC))
}
