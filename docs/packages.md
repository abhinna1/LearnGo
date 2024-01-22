# Packages

A go module can be defined as a package with the help of the go.mod file.
Here, we can specify the module name and the require Go version for the package/project.

```go
module "github.com/abhinna1/LearnGo"

go 1.21
```

We can define multiple other packages in the project using the 'package' keyword in Go.

```go
package database
```

## Importing packages

We can import a go package using the import function.
'fmt' is an in-built package in Go that provides various functions.

```go
import "fmt"
```

We can also import user-created packages.
```go
import "github.com/abhinna1/LearnGo/database"
```
