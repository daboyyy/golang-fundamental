package main

func main() {
	c, x, b := addReturnTuple(10, 20)
	println(c, x, b)
	sum, _, _ := addReturnTuple(10, 20)
	println(sum)
	println("")

	calc(add)
	calc(sub)
	// in javascript it's call closure
	calc(func(a, b int) int {
		return a * b
	})
	println("")

	arr := []int{1, 2, 3, 4, 5}
	println(sumArr(arr))
	println(sumManyInt(1,2,3,4,5))
}

// basic function
func add(a int, b int) int {
	return a + b
}
func addShortTypeDeclaration(a, b int) int {
	return a + b
}
// tuple
func addReturnTuple(a, b int) (int, string, bool) {
	return a + b, "ok", true
}

func sub(a, b int) int {
	return a - b;
}

// function params
func calc(f func(int, int)int) {
	sum := f(50, 10)
	println(sum)
}
// end basic function

// array function
func sumArr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumManyInt(arr ...int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
// end array function
