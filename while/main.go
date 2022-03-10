package main

// go doesn't have DoWhile

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	i := 0
	for i < len(arr) {
		println(arr[i])
		i++
	}
}
