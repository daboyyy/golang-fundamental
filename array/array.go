package main

import "fmt"

func main() {
	var arr [3]int;
	fmt.Printf("arr %v \n", arr)
	fmt.Printf("arr %#v \n", arr)
	println("")

	var arr2 [3]int = [3]int{1,2,3};
	fmt.Printf("arr2 %v \n", arr2)
	arr3 := [3]int{4,5,6};
	fmt.Printf("arr3 %v \n", arr3)
	arr3[0] = 40
	fmt.Printf("new arr3 %v \n", arr3)
	println("")

	twoDimensionArr := [2][3]int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Printf("twoDimensionArr %v \n", twoDimensionArr)
	println("")

	dynicmicLengthArr := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("dynicmicLengthArr %#v \n", dynicmicLengthArr)
	println("")
}
