package main

type Person struct {
	Name string // if you rename this to name you can call only in this package can't call from outside (package scope)
	Age int
}

func HelloFunc(p Person) string {
	return "Hello " + p.Name
}

func (p Person) HelloMethod() string {
	return "Hello " + p.Name
}

func main() {
	person := Person{
		Name: "Bond",
		Age: 18,
	}
	println(person.Name)
	println(person.Age)
	println("")

	// function
	println(HelloFunc(person))
	// method
	println(person.HelloMethod())
}
