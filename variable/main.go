package main

import "fmt"

// package scope (global variable)
var global int
var global2 = 10

func main() {
	var variable int
	_ = variable // ignore a can't compile if block scope doesn't use

	var variable2 int
	variable2 = 10
	fmt.Printf("Printout variable2 = %v \n", variable2)

	variable3 := 20
	fmt.Printf("Printout variable3 = %v \n", variable3)
	println("")

	// every variable has zero value (default value)
	var zeroValue int // default value = 0
	fmt.Printf("default value of zeroValue = %v \n", zeroValue)
	var zeroValue2 string // default value = ""
	fmt.Printf("default value of zeroValue2 = %v \n", zeroValue2)
	var zeroValue3 bool // default value = false
	fmt.Printf("default value of zeroValue3 = %v \n", zeroValue3)
}
