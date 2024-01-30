# Scope

The scope of an attribute defines from where it can be accessed.

## Local scope

A variable declared inside a function can only be accessed within that function.
The memory location storing the value in that variable is wiped once the function is done executing.
Therefore, a local varible defined inside a function cannot be accessed from outside the scope of that function.

```go
func FunctionA() {
    var number1 int = 1
}
```

Here, the variable 'number1' is a local variable and cannot be accessed outside of the scope function 'FunctionA'.

## Global score

A variable declared at a global scope, or outside a function can be accessed throughout the file.

```go
import "fmt"

const GLOBALVARIABLE int = 0

func main(){
    fmt.Println(GLOBALVARIABLE)
}
```

Here, the variable 'GLOBALVARIABLE' is defined at a global scope, therefore, it can be accessed from anywhere in that file, including from inside the functions.

## Package scope

Similar to the scope of variables and functions in a file, their scopes can also differ when trying to access them from different packages.

In GO, when we create a variable in a package, in order to make them public (accessible from another package) we need to name them with a capital initial. And for a private variable or function, the naming should start with a lowe-case initial.

```go

package database

import "fmt"

const dbName string = "database1"

func ConnectDb(){
    fmt.Println("Connected to database.")
}

```

In the above code example, the variable 'dbName' is only accessible from within the database package as it's naming begins with a lowercase.
The 'ConnectDb' function however, is public and can be accessible from different packages as the naming starts with an upper case.
