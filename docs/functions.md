# Functions

Functions in Go, are like in any other programming languages. They are just a block of code that you can execute by calling them anywhere in your program, reducing repetition of code.

## [Code Example](../tutorials/functions/main.go)

## Defining a function

In go, we use the `func` keyword to define a function.

```go
func helloWorld() {
    // Function that prints hello world.

    fmt.Println("Hello world!")
}
```

## Arguments and parameters

As per the definition, parameters are simply the values that a function asks for when they are called. And arguments are the values that you provide to the parameters when the function is called.

A function can take multiple parameters. In go, you should also specify the type of the parameters while defining the function.

In the `getSum` function defined below, num1 and num2 are two integer parameters required by the function. It is necessary to provide these parameters while calling the `getSum` function.

```go
func getSum(num1 int, num2 int) int {
    // Function that takes two integer arguments
    // and returns their sum.

    return num1 + num2
}
```

## Calling a function

You can call a function simply with the name of the function followed by `()`.
If a function takes in any arguments, you should provide them within the parenthesis. The type of arguments should be the same as defined by the function in order to avoid any runtime errors.

```go
// Calling a function that doesn't require arguments.
helloWorld();

// Calling a function that requires arguments.
var n1 int = 1;
var n2 int = 3;

var sum int = getSum(n1, n2);
```

## main function

The main function is the entrypoint of a Go program. Meaning, it is the function that is first executed when running your program.

**Note: Your main function should be inside the main package.**

```go
package main

import "fmt"

func main() {
    // The entry-point of the program.

    fmt.Println("Hello World!")
    helloWorld()
    var n1 int = 1;
    var n2 int = 3;

    var sum int = getSum(n1, n2);
    fmt.Println(sum)
}
```
