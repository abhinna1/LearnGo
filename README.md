# LearnGo

Open-Source public repository to learn GO.

## Content

1. [Variables 🗃️](docs/variables.md)
2. [Conditional Statements 🤔](docs/conditional.md)
3. [Functions 📲](docs/functions.md)
4. [Packages 📦](docs/packages.md)
5. [Scope 🔭](docs/scope.md)

## Code examples

Code examples for different topics are available inside the [tutorials package](tutorials/).
These examples cover various titles revolving the syntax and concepts in Go.

## Tests

Unit-tests are a vital part of any project. Each package in this project contains one or multiple test files with names in the format `*_test.go` which have the unit-tests for all code examples and packages in the project.

The tests can be debugged in VS-Code or through the cli command `go test ./tutorials/...`.

This command will run all the tests available in the tests module.
However, if you want to run any particular test file, you can run `go test ./tutorials/package/filename_test.go`

You can also add other arguments along with the tests.

- `go test ./tests -v`: additional details of the tests.
- `go test ./tests -cover`: coverage of the tests.

## Coverage

Test coverage is an essential part of this project to ensure the code quality.
You can view the coverage of the test cases by adding a `-cover` argument while running the test cases.

Additionally, you can also generate a coverage report file with the `-coverprofile ./coverage.out` argument.
This creates a `coverage.out` file in the project root. You can then use this file to visually represent the test coverage in the web-browser with the command `go tool cover -html=coverage.out`.

`
  NOTE: The coverage report should be ignored if you want to contribute to the project so make sure to strictly name your report file as coverage.out as it is included in the gitignore file. Additionally, it is mandatory for all the packages and contributions to have a 100% code coverage for a successful contribution.
`
