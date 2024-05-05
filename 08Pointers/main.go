package _8Pointers

import "fmt"

func sumValues(num1 int, num2 int) int {
	var c int = num1 + num2
	num1 = 0
	fmt.Printf("Value of a inside functions: %v\n", num1)
	return c
}

func sumValues2(num1 *int, num2 *int) int {
	var c int = *num1 + *num2
	*num1 = 0
	fmt.Printf("Value of a inside functions: %v\n", *num1)
	return c
}

func main() {
	/*
		 Pointers:
			 - A pointer is a variable that stores the memory address of another variable.
			 - A pointer is used to store the address of a variable.
			 - A pointer is used to access the value of a variable indirectly.
			 - A pointer is used to pass the address of a variable to a function.
			 - A pointer provide ways to change the value located on a memory location.
			 - A pointer is used to allocate memory dynamically.
		& operator:
			- The & operator is used to get the address of a variable.
			- It is known as the address-of operator.
			- Syntax: &variableName
		* operator:
			- The * operator is used to get the value of a variable through a pointer.
			- It is known as the dereference operator.
			- Syntax: *pointer
		To declare a pointer:
			- var pointerName *dataType
			- Example: var ptr *int
			- Zero value of a pointer is nill
		To assign a pointer:
			- pointerName = &variableName
			- Example: ptr = &a
		To access the value of a pointer:
			- *pointerName
			- Example: *ptr
	*/

	var a int = 10
	var b *int
	b = &a

	println("Value of a is: ", a)
	println("Address of a (&a) is: ", &a)
	println("Value of pointer b is: ", b)
	println("Value stored on the address of *b is: ", *b)

	fmt.Printf("Type and value of a : %T, %v \n", a, a)
	fmt.Printf("Type and value of &a : %T, %v \n", &a, &a)
	fmt.Printf("Type and value of b : %T, %v \n", b, b)
	fmt.Printf("Type and value of *b : %T, %v \n", *b, *b)

	name := "Aldrich"
	ptr_name := &name
	fmt.Printf("Type and value of a : %T, %v \n", name, name)
	fmt.Printf("Type and value of a : %T, %v \n", ptr_name, ptr_name)
	fmt.Printf("Type and value of a : %T, %v \n", *ptr_name, *ptr_name)

	/*
		Passing by value in functions:
			- Function is called by passing the value stored in the variable address.
			- Param is stored into another location of your memory
			- When variable is modified, only copy is affected
	*/

	println("\n##### Example callin by value #####")
	var num1, num2, num3 int
	num1 = 6
	num2 = 23
	num3 = sumValues(num1, num2)
	fmt.Printf("Value of num1 before: %v\n", num1)
	fmt.Printf("Result %v\n", num3)
	fmt.Printf("Value of num1 outside functions: %v\n", num1)

	/*
		Passing by reference in Functions
			- The address of the variable is passed into function call as param.
			- All operations in function are performed on the value of the address.
	*/

	println("\n##### Example callin by reference #####")
	num3 = sumValues2(&num1, &num2)
	fmt.Printf("Value of num1 before: %v\n", num1)
	fmt.Printf("Result %v\n", num3)
	fmt.Printf("Value of num1 outside functions: %v\n", num1)
}
