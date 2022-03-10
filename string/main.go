package main

import "fmt"
import "unicode/utf8"

func main() {
	engString := len("a")
	fmt.Printf("engString length = %v \n", engString)

	numString := len("1")
	fmt.Printf("numString length = %v \n", numString)

	thString := len("ก")
	fmt.Printf("thString length = %v \n", thString)
	println("")

	stringLength := utf8.RuneCountInString("ก")
	fmt.Printf("stringLength length = %v \n", stringLength)
}
