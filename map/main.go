package main

import "fmt"

// in python it's call Dictionary

func main() {
	//				    key	 value
	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["us"] = "United State"

	fmt.Printf("countries th = %v \n", countries["th"])
	fmt.Printf("countries us = %v \n", countries["us"])
}
