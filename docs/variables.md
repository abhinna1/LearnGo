
# Variables

## Declaring variables

### [Code Example](../tutorials/variables/variables.go)

Variables can be defined in Go using the var keyword followed by the variable name and type respectively.

```go
var num1 int
```

### Variable short declaration

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```go
name := "Abhinna"
username, password, isOnline := "abhinna1", "@1234", true
```

const can be defined using the **const** keyword.

```go
const THRESHOLD = 3
```
