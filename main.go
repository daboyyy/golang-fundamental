package main

// external
import "fmt"
// internal
import "golang-fundamental/packages"
import "golang-fundamental/oop"

func main() {
	// package
	println(packages.Status())
	println(packages.CustomerName)
	println("")

	// strcut-oop
	person := oop.Person{}
	person.SetName("Bond")
	person.SetAge(18)
	fmt.Printf("%v \n", person.GetName())
	fmt.Printf("%v \n", person.GetAge())
}
