package main

// in JavaScript it's call ForEach

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}

	for i, v := range arr {
		println(i, v)
	}
	println("")

	for _, v := range arr {
		println(v)
	}
}
