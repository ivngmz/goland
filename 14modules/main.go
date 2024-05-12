package _4modules

import (
	"fmt"
	"github.com/ivngmz/golang/14modules/map"
	"github.com/ivngmz/golang/14modules/reduce"
)

func main() {
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
	result := map.MapIntToInt(f, a)

	fmt.Println(result)

	// Call ReduceIntToInt with the function and slice as arguments
	reducedResult := reduce.ReduceIntToInt(r, a)

	fmt.Println(reducedResult)

}
