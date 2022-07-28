package main

import "fmt"

func min(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	val := n[0]
	for _, v := range n {
		if v < val {
			val = v
		}
	}
	return val
}

func main() {
	num := min(3, 1, 6, 8, -2, 65)
	fmt.Printf("The mininum is %d\n", num)
	slice := []int{3, 8, 1, 9, 3, 66, 22, 8}
	num = min(slice...)
	fmt.Printf("The minimum in the slice is %d\n", num)
}
