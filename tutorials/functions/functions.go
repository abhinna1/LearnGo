package main

import "fmt"

func helloWorld() {
	// Function that prints hello world.

	fmt.Println("Hello world!")
}

func getSum(num1 int, num2 int) int {
	// Function that takes two integer arguments
	// and returns their sum.

	return num1 + num2
}

func main() {
	// The entry-point of the program.

	fmt.Println("Hello World!")
	helloWorld()
	var n1 int = 1;
	var n2 int = 3;

	var sum int = getSum(n1, n2);
	fmt.Println(sum)
}
