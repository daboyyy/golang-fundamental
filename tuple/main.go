package main

import "fmt"

// can return value more then one

func main() {
	x, y := 10, 20
	fmt.Printf("x value = %v \n", x)
	fmt.Printf("y value = %v \n", y)
	println("")

	// tuple with map
	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["us"] = "United State"

	countryData, ok := countries["jp"]
	if ok {
		println(countryData)
	} else {
		println("No Value")
	}
}
