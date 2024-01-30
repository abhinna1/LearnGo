# Conditional Statements

Conditional statements allow developers to control the flow of a program based on certain conditions.
Go provides the standard `if`, `else`, and `switch` statements for handling conditions in a concise and expressive manner.

[***Code Example***](../tutorials/conditional/conditional.go)

## if, elseif, else

If, else if and else are the most straight-forward way of implementing conditionals in Go.
This can be achieved by the following syntax.

```go
func AccessControlSwitchCases(role string) int {
    if role == "owner" {
        return 1
    } else if role == "admin" {
        return 2
    } else if role == "user" {
        return 3
    }
    return 0
}
```

## Switch case

Switch statement evaluates an expression and executes the code block associated with the first matching case,
offering a concise way to handle multiple conditions without requiring explicit break statements between cases.
It is similar to using if-else. However, the values provided in the switch case is compared directly with the possible
cases.

```go
func AccessControl(role string) int {
    switch role {
    case "owner":
        return 1
    case "admin":
        return 2
    case "user":
        return 3
    }
    return 0
}
```
