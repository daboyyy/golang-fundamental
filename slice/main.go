package main

import "fmt"

// in python it's call List

func main() {
	slice := []int{1,2,3}
	fmt.Printf("slice value = %#v \n", slice)
	slice = append(slice, 4)
	fmt.Printf("append slice value = %#v \n", slice)
	println("")

	sliceLength := len(slice)
	fmt.Printf("slice length = %v \n", sliceLength)
	println("")

	//				0	1	2	3	4	5	6	7	8
	slice2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	cloneSlice2 := slice2[1:]
	fmt.Printf("cloneSlice2 value = %v \n", cloneSlice2)
}
