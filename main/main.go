package main

import "fmt"

/*markdown explanation
# Data Types
- String: A string is a sequence of characters enclosed in double quotes
- Number:
	- Integer: Whole number
	- Float: Decimal number
- Boolean: True or False
- Array: Fixed size list
- Slice: Flexibel array
- Map: Key-value pair

# What are Data Types?
- Data types are the classification or categorization of data items.
- Describe the operations that can be performed on the data.
- Define the way the data will be stored in memory.

# Memory Allocation
- Memory is allocated based on the data type.
- The size of the data type determines the amount of memory allocated.
- The data type also determines the range of values that can be stored.
- Memory allocation based on every kind of data:
	- String: 2 bytes per character
	- Number:
		- Integer: 8 bytes
		- Float: 8 bytes
	- Boolean: 1 byte
	- Array: 8 bytes per element
	- Slice: 24 bytes
	- Map: 8 bytes per key-value pair

# Static vs Dynamic Typing
- Static Typing: The data type is determined at compile time.
	- Compiler checks the data type and throws an error if the data
	type is not compatible.
	- Go is a statically typed language.
- Dynamic Typing: The data type is determined at runtime.
	- The data type is checked at runtime.
	- Python is a dynamically typed language.
- Adventages of Static Typing:
	- Faster execution.
	- Bugs can be often caught at compile time.
	- Better data integrity.
- Adventages of Dynamic Typing:
	- Faster to write code.
	- More flexibility.
	- Easier to learn.
- Goland:
	- Has a concept of type inference, which means that the compiler
	can infer the type of a variable based on the value assigned to it.
	- It is a fast statically typed language that feels like a
	dynamically typed language.
*/

func main() {
	var num1, num2, result int
	fmt.Println("Enter first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter second number: ")
	fmt.Scanf("%d", &num2)
	result = num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, result)
}
