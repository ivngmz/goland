package main

import (
	"fmt"
	mapreduce "github.com/ivngmz/golang/14modules"
)

func main() {
	/* Library management:
	Base concepts to manage libraries:
		- repository -> where sourcecode project is stored
		- packages -> a collection of source files in the same directory that are compiled together
		- modules -> a collection of related Go packages that are versioned together as a single unit

	Go modules:
	- is the root of dependency management in Go.
	- is a collection of related Go packages that are versioned together as a single unit.
	- a module project is defined by a go.mod file and must be located in the root of the project.
	- the module specifier is the path to the module root directory.
	- some commmands:
		- could be created by running go mod init <module-name>.
		- could be updated by running go get <module-name>.
		- could be removed by running go mod tidy.
		- could be downloaded by running go mod download.
		- could be checked by running go mod verify.
		- could be updated by running go mod vendor.
		- could be updated by running go mod graph.
		- could be updated by running go mod why.
		- could be updated by running go mod edit.
	- module must be public, so it must be stored in a public repository.
	*/

	// Define a function to be passed to MapIntToInt
	f := func(i int) int {
		return i * 2
	}

	// Define a function to be passed to ReduceIntToInt
	r := func(x, y int) int {
		return x + y
	}

	// Define a slice of integers
	a := []int{1, 2, 3, 4, 5}

	// Call MapIntToInt with the function and slice as arguments
	result := mapreduce.map.mapIntToInt(f, a)

	fmt.Println(result)

	// Call ReduceIntToInt with the function and slice as arguments
	reducedResult := mapreduce.reduce.reduceIntToInt(r, a)

	fmt.Println(reducedResult)

}
