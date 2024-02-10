package functions

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
