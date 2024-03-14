package sample

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func main() {
	fmt.Sprintf("Add 5 + 3 = %v", add(5, 3))
	fmt.Sprintf("Subtract 5 - 3 = %v", subtract(5, 3))
}
