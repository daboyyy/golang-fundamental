package main

import "fmt"

func main() {
	point := 51
	if point >= 30 {
		fmt.Println("Point more than 30 !!!")
	} else if point < 100 && point > 50 {
		fmt.Println("Point less than 100 and more than 50 !!!")
	} else {
		fmt.Println("Point less than 31 !!!")
	}
}
