package _3userInput

import (
	"fmt"
	"strconv"
)

func main() {
	const pi float64 = 3.14159265359
	var num1, num2 string
	var numA, numB int
	var result int

	/*
		- PrintF is a function that is used to format the output.
		- The & operator is used to get the memory address of a variable.
		- Options:
			- %v: Value in a default format.
			- %s: Value as a string.
			- %d: Value as a decimal integer.
		- count 1 and err1 are used to capture the number of arguments and errors.
	*/
	fmt.Printf("Enter a number:  \n")
	count1, err1 := fmt.Scanf("%s", &num1)
	fmt.Printf("Enter another number:  \n")
	count2, err2 := fmt.Scanf("%s", &num2)
	fmt.Printf("\n--- count1: %d, err1: %v\n", count1, err1)
	fmt.Printf("\n--- count1: %d, err1: %v\n", count2, err2)

	/*
		- The type of num1 is string.
	*/
	fmt.Printf("\n--- The type of %v is %T: ", num1, num1)
	fmt.Printf("\n--- The type of %v is %T: ", num2, num2)

	/*
		- Functions on the strconv package:
			- Atoi is a function that is used to convert a string to an integer.
			- Itoa is a function that is used to convert an integer to a string.
			- ParseBool is a function that is used to convert a string to a boolean.
			- ParseFloat is a function that is used to convert a string to a float.
			- ParseInt is a function that is used to convert a string to an integer.
			- ParseUint is a function that is used to convert a string to an unsigned integer.
			- FormatBool is a function that is used to convert a boolean to a string.
			- FormatFloat is a function that is used to convert a float to a string.
		- Another option is to cast the string to an integer.
			- Example:
				- numA = int(num1)


	*/
	numA, errA := strconv.Atoi(num1)
	fmt.Printf("errA: %v\n", errA)
	numB, errB := strconv.Atoi(num2)
	fmt.Printf("errB: %v\n", errB)
	result = numA + numB
	fmt.Println("\nThe sum of the two numbers is: ", result)
	fmt.Printf("\n--- The type of %v is %T: ", num1, num1)
	/*
		- const was used to declare a constant.
			- scope: The scope of a constant is the region of the program where the constant is accessible.
			- Constants are accessible throughout the program.
			- Constants are immutable.

	*/
	fmt.Printf("\n--- The type of contant %v is %T: and value is: %f", pi, pi, pi)
}
