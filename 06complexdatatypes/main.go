package _6complexdatatypes

import "fmt"

func fibonacci(arrayFibonacci *[50]int) {
	arrayFibonacci[0] = 0
	arrayFibonacci[1] = 1
	for i := 2; i < 50; i++ {
		arrayFibonacci[i] = arrayFibonacci[i-1] + arrayFibonacci[i-2]
	}
}

func main() {
	/* Arrays
	- An array is a collection of elements of the same type.
	- The size of an array is fixed.
	- The elements of an array are stored in contiguous memory locations.
	- The elements of an array can be accessed using an index.
	- The index of an array starts from 0.
	- The syntax for declaring an array is:
		- var arrayName [size]dataType
	- Example:
		- var numbers [5]int
	*/

	// Calls function to calculate fibonacci series
	var arrayFibonacci [50]int
	fibonacci(&arrayFibonacci)
	fmt.Printf("The fibonacci series for 50 first elements is: \n")
	fmt.Println(arrayFibonacci)
	fmt.Println("\n-> Value of index 5 of array is: $d", arrayFibonacci[5])

	/* Slices
	- A slice is a reference to an underlying array.
	- The size of a slice is not fixed.
	- The elements of a slice are stored in contiguous memory locations.
	- The elements of a slice can be accessed using an index.
	- The index of a slice starts from 0.
	- The syntax for declaring a slice is:
		- var sliceName []dataType
	- Example:
		- var numbers []int
	*/

	/* Maps
	- A map is a collection of key-value pairs.
	- The keys of a map are unique.
	- The keys of a map are of the same type.
	- The values of a map are of the same type.
	- The elements of a map are not stored in contiguous memory locations.
	- The elements of a map can be accessed using a key.
	- The syntax for declaring a map is:
		- var mapName map[keyType]valueType
	- Example:
		- var cities map[string]string
	*/

	/* Structs
	- A struct is a collection of fields.
	- The fields of a struct can be of different types.
	- The fields of a struct are stored in contiguous memory locations.
	- The fields of a struct can be accessed using a dot operator.
	- The syntax for declaring a struct is:
		- type structName struct {
			field1 dataType1
			field2 dataType2
	}
	- Example:
		- type Person struct {
			name string
			age int
	}

	*/
}
