package main

import "fmt"

// Function that prints hello world.
func helloWorld() {
	fmt.Println("Hello world!")
}

// Function to get sum of two numbers.
// 	Params:
// 		num1(int): first number
// 		num2(int): second number
// 	Returns:
// 		int: sum of num1 and num2
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

// The entry-point of the program.
func main() {
	fmt.Println("Hello World!")
	helloWorld()
	var n1 int = 1
	var n2 int = 3

	var sum int = getSum(n1, n2)
	fmt.Println(sum)
}
