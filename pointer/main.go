package main

import "fmt"

func main() {
	x1 := 10
	y1 := x1
	fmt.Println(&x1)
	fmt.Println(&y1)
	println("")

	x2 := 20
	y2 := &x2 // same as var y2 *int
	fmt.Println(&x2)
	fmt.Println(y2)
	println("")

	println("y2 =", *y2)
	x2 = 10
	fmt.Println("x2 = 10 then y2 =", *y2)
	println("")

	var output int
	sumViaPointer(&output)
	println(output)
}

func sumViaPointer(result *int) {
	a := 10
	b := 20
	*result = a + b
}
