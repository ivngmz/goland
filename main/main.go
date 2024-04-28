package main

import "fmt"

func fibonacci(arrayFibonacci *[50]int) {
	arrayFibonacci[0] = 0
	arrayFibonacci[1] = 1
	// using len(arrayFibonacci) to get the length of the array
	for i := 2; i < len(arrayFibonacci); i++ {
		arrayFibonacci[i] = arrayFibonacci[i-1] + arrayFibonacci[i-2]
	}
	// looping using range of the array
	for i, value := range arrayFibonacci {
		fmt.Printf("Index %d: %d\n", i, value)
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
	- make function can be used to create an array. (capacity is optional, default is size)
		- Syntax: arrayName := make([]dataType, size, [capacity])
	- Example:
		- var numbers [5]int
	 	- var numbers := make([]int, 5, 10)
	*/

	// Calls function to calculate fibonacci series
	var arrayFibonacci [50]int
	fibonacci(&arrayFibonacci)
	fmt.Printf("The fibonacci series for 50 first elements is: \n")
	fmt.Println(arrayFibonacci)
	fmt.Printf("\n-> Value of index 5 of array is: %d",
		arrayFibonacci[5])

	// Array with assigned values
	var arrayAsignedValues = [5]int{1, 2, 3, 4, 5}
	println("\n-> Value of index 2 of arrayAsignedValues is: ",
		arrayAsignedValues[2])

	// Array with := operator
	arrayWithOperator := [5]int{1, 2, 3, 4, 5}
	println("\n-> Value of index 3 of arrayWithOperator is: ",
		arrayWithOperator[3])

	// Array using elipsis
	arrayUsingElipsis := [...]int{1, 2, 3, 4, 5}
	println("\n-> Value of index 4 of arrayUsingElipsis is: ",
		arrayUsingElipsis[4])
	// Change value of index 4
	arrayUsingElipsis[4] = 10
	println("\n-> Value of index 4 has been changed on arrayUsingElipsis, new value is: ",
		arrayUsingElipsis[4])

	// Multidimensional array example
	multidimensionalArray := [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("\n-> Value of index [%d][%d] of multidimensionalArray is: %d",
				i, j, multidimensionalArray[i][j])
		}
	}

	/* Slices
	- A slice is a reference to an underlying array.
	- The size of a slice is not fixed. I is more flexible than an array.
	- The elements of a slice are stored in contiguous memory locations.
	- The elements of a slice can be accessed using an index.
	- The index of a slice starts from 0.
	- internally slice have a pointer to the first element of the array,
		the length of the slice and
		the capacity of the slice.
	- The syntax for declaring a slice is:
		- var sliceName []dataType
	- Example:
		- var numbers []int
	*/

	// Slice with assigned values
	sliceAssignedValues := []int{1, 2, 3, 4, 5}
	fmt.Printf("\n-> Value of index 2 of sliceAssignedValues is: %d", sliceAssignedValues[2])

	// Slice with := operator
	sliceWithOperator := []int{1, 2, 3, 4, 5}
	println("\n-> Value of index 3 of sliceWithOperator is: ", sliceWithOperator[3])

	// Slice using elipsis
	sliceUsingElipsis := []int{1, 2, 3, 4, 5}
	println("\n-> Value of index 4 of sliceUsingElipsis is: ", sliceUsingElipsis[4])

	// Slice using make
	sliceUsingMake := make([]int, 5)
	println("\n-> Value of index 4 of sliceUsingMake is: ", sliceUsingMake[4])

	// Slice using make with capacity and length
	sliceUsingMakeWithCapacityAndLength := make([]int, 5, 10)

	// Slice using make with capacity and length
	sliceUsingMakeWithCapacityAndLength[0] = 1

	// Slice using append
	sliceUsingAppend := []int{1, 2, 3, 4, 5}
	fmt.Printf("\n-> Initial values of sliceUsingAppend are: %d", sliceUsingAppend)
	sliceUsingAppend = append(sliceUsingAppend, 6, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("\n-> Final values of sliceUsingAppend are: %d", sliceUsingAppend)

	// delete elements from slice
	sliceUsingAppend = append(sliceUsingAppend[:2], sliceUsingAppend[7:]...)
	fmt.Printf("\n-> Final values of sliceUsingAppend after deleting element at index 2 to 6 are: %d",
		sliceUsingAppend)

	// Slice using copy
	sliceUsingCopy := []int{1, 2, 3, 4, 5}
	sliceUsingCopy2 := make([]int, 5)
	copy(sliceUsingCopy2, sliceUsingCopy)

	// Slice using range
	sliceUsingRange := []int{1, 2, 3, 4, 5}
	for i, value := range sliceUsingRange {
		fmt.Printf("\n-> Value of index %d of sliceUsingRange is: %d", i, value)
	}

	/* Maps
	- A map is a collection of key-value pairs.
	- The zero value of a map is nil and do not contain any key-value pairs.
	- The keys of a map are unique.
	- The keys of a map are of the same type.
	- The values of a map are of the same type.
	- The elements of a map are not stored in contiguous memory locations.
	- The elements of a map can be accessed using a key.
	- Elements of a map are mutable and capacity is not fixed.
	- The syntax for declaring a map is:
		- var mapName map[keyType]valueType[{key1: value1, key2: value2}]
	- Example:
		- var cities map[string]string
		- var cities map[string]string{"Indonesia": "Jakarta", "Malaysia": "Kuala Lumpur"}
	*/

	// Map with assigned values
	mapAssignedValues := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("\n-> Value of key 'two' of mapAssignedValues is: %d", mapAssignedValues["two"])

	// Map using make
	mapUsingMake := make(map[string]int)
	mapUsingMake["one"] = 1
	mapUsingMake["two"] = 2
	mapUsingMake["three"] = 3

	// Map using range
	mapUsingRange := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range mapUsingRange {
		fmt.Printf("\n-> Value of key %s of mapUsingRange is: %d", key, value)
	}

	// Getting value from map
	arrayKey := []string{"two", "four"}
	/*
		- Using blank identifier to ignore the first return value of the map.
		- The blank identifier is represented by an underscore (_).
		- The blank identifier is used to ignore the return value of a function.
	*/
	for _, key := range arrayKey {
		value, ok := mapUsingRange[key]
		if ok {
			fmt.Printf("\n-> Key %s found in mapUsingRange", key)
			fmt.Printf("\n-> Value of key 'two' of mapUsingRange is: %d", value)
		} else {
			fmt.Printf("\n-> Key %s NOT found in mapUsingRange", key)
		}
	}

	// print map
	fmt.Printf("\n-> mapUsingRange: %v", mapUsingRange)

	// Adding value to map
	mapUsingRange["four"] = 4
	fmt.Printf("\n-> Value of key 'four' of mapUsingRange is: %d", mapUsingRange["four"])

	// Update value in map
	mapUsingRange["four"] = 5
	fmt.Printf("\n-> Value of key 'four' of mapUsingRange after updating is: %d", mapUsingRange["four"])

	// print map
	fmt.Printf("\n-> Previous map is: %v", mapUsingRange)

	// Delete value in map
	delete(mapUsingRange, "four")
	fmt.Printf("\n-> Value of key 'four' of mapUsingRange after deleting is: %d", mapUsingRange["four"])

	// print map again
	fmt.Printf("\n-> Later map is: %v", mapUsingRange)

	// iteration over map
	for key, value := range mapUsingRange {
		fmt.Printf("\n-> Key: %s, Value: %d", key, value)
	}

	/* truncate map: two options:
	- reinitialize the map
		- mapUsingRange = map[string]int{}
	- delete all elements by looping over the map
		- see it below
	*/
	for key := range mapUsingRange {
		delete(mapUsingRange, key)
	}

}
