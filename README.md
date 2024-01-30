# LearnGo

Open-Source public repository to learn GO.

## Content

1. [Scope 🔭](docs/scope.md)
2. [Variables 🗃️](docs/variables.md)
3. [Packages 📦](docs/packages.md)

## Code examples

Code examples for different topics are available inside the [tutorials package](tutorials/).
These examples cover various titles revolving the syntax and concepts in Go.

## Tests

Unit-tests are a vital part of any project. The [tests module](tests/) contain unit-tests for
all code examples and packages in the project.

The tests can be debugged in VS-Code or through the cli command `go test ./tests`.

This command will run all the tests available in the tests module.
However, if you want to run any particular test file, you can run `go test ./tests/filename_test.go`

You can also add other arguments along with the tests.

- `go test ./tests -v`: additional details of the tests.
- `go test ./tests -cover`: coverage of the tests.
