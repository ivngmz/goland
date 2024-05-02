package _7functions

import "fmt"

func printMsg() {
	println("Hello, World!")
}
func logMsg(msg string) string {
	return msg
}

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func blankIdentifier() (int, int) {
	return 1, 2
}

func sumVar(nums ...int) int {
	total := 0
	/*Blank Identifier:
	- The blank identifier (_) is used to discard values.
	- The blank identifier is used to discard values that are not needed.
	- The blank identifier is represented by an underscore (_).
	*/
	numericDate, _ := blankIdentifier()
	println("This is the return of the blank identifier function example:")
	println(numericDate)

	for _, num := range nums {
		total += num
	}
	return total
}

func recursiveFunction(n int) int {
	// Example calling this function:
	// - sum := recursiveFunction(5)
	if n == 0 {
		return 0
	}
	return n + recursiveFunction(n-1)
}
func anonymousFunction() {

}
func recipe(principal string, ingredients ...string) {
	var allIngredients []string
	allIngredients = append(allIngredients, principal)
	allIngredients = append(allIngredients, ingredients...)
	println(allIngredients)
}

func main() {
	// Functions are used to:
	// - Encapsulate code
	// - Make code reusable
	// - Make code easier to read
	// - Make code easier to maintain

	// Naming conventions for functions:
	// - Use camelCase for function names
	// - Use descriptive names for function names
	// - Use verbs for function names
	// - Use nouns for argument names
	// - Use descriptive names for argument names
	// - Use lower case for function names
	// - Use lower case for argument names
	// - Use lower case for return types
	// - Use lower case for local variables
	// - Use lower case for package names

	// Functions are defined with the "func" keyword
	// The function name is "main" and it is in the "main" package and is called when the program is run

	// Functions are defined with the "func" keyword
	// Functions can have arguments and return types:
	// - Arguments are values passed to the function
	// - Return types are values returned by the function
	// - Functions can have multiple arguments and return types
	// - Arguments and return types are optional
	// Example of functions with arguments and return types:
	// - func add(a int, b int) int { return a + b }
	// Example calling a function with arguments and return types:
	// - sum := add(1, 2)
	// Example of functions with no arguments and no return types:
	// - func hello() { fmt.Println("Hello") }
	// Example calling a function with no arguments and no return types:
	// - hello()
	var log string
	log = logMsg("Hello, World!")
	println(log)
	printMsg()

	var mod, rest int = divide(10, 3)
	println("Module=", mod)
	println("Rest=", rest)

	// Variadic functions:
	// - Variadic functions take a variable number of arguments
	// - The arguments are passed as a slice
	// - The arguments are accessed using the "..." operator
	// - The "..." operator is used to pass a slice to a variadic function
	// Example of a variadic function:
	// - func sum(nums ...int) int { total := 0; for _, num := range nums { total += num }; return total }
	// Example calling a variadic function:
	// - total := sum(1, 2, 3, 4, 5)
	// - total := sum([]int{1, 2, 3, 4, 5}...)
	// - total := sum(nums...)

	var result int
	result = sumVar(1, 2, 3, 4, 5)
	println(result)

	var ingredients = []string{"flour", "sugar", "eggs", "butter"}
	var principal = "cake"
	var allIngredients []string
	allIngredients = append(ingredients, principal)
	for _, ingredient := range allIngredients {
		println(ingredient)
	}

	/*
		- Recursion is a technique in which a function calls itself.
		- Recursion is used to solve problems that can be broken down into smaller problems.
		- Recursion is used to solve problems that can be solved by solving smaller instances of the same problem.
		- Recursion is used to solve problems that can be solved by solving smaller instances of the same problem.
	*/
	totalSum := recursiveFunction(5)
	println("The total sum is: ", totalSum)

	// Anonymous functions:
	// - Anonymous functions are functions without a name.
	// - Anonymous functions are used to define a function inline.
	// - Anonymous functions are used to define a function inside another

	anonymousMessage := func() { println("Hello, World!") }
	fmt.Printf("%T\n", anonymousMessage)
	anonymousMessage()
	anonymousResult := func(a int, b int) {
		modCalc := a / b
		println(modCalc)
	}
	fmt.Printf("%T\n", anonymousResult)
	anonymousResult(34243, 32)

	// Higher-order functions:
	// - Higher-order functions are functions that take other functions as arguments.
	// - Higher-order functions are functions that return other functions.
	// - Higher-order functions are used to create abstractions.
	// - Higher-order functions are used to create reusable code.
	// - Higher-order functions are used to create modular code.
	// - Higher-order functions are used to create composable code.
	// - Higher-order functions are used to create flexible code.
	// - Higher-order functions are used to create extensible code.
	// Example of a higher-order function:
	// - func apply(f func(int, int) int, a int, b int) int { return f(a, b) }
	// Example calling a higher-order function:
	// - result := apply(add, 1, 2)

	// TODO: Check problem with the higher-order function
	//func apply(f func(int, int) int, a int, b int) int { return f(a, b) }
	//sumResult := sumDiff(1,2, sum)
	//diffResult := sumDiff(1,2, diff)
	//println("Higher fucntion result:")
	//println(sumResult)
	//println(diffResult)

	// Deffered functions:
	// - Deferred functions are executed after the surrounding function returns.
	// - Deferred functions are executed in the reverse order of their calls.
	// - Deferred functions are used to clean up resources.
	// - Deferred functions are used to release locks.
	// - Deferred functions are used to close files.
	// - Deferred functions are used to free memory.
	// - Deferred functions are used to log errors.
	// - Deferred functions are used to recover from panics.
	// Example of a deferred function:
	// - func cleanup() { fmt.Println("Cleaning up...") }
	// - defer cleanup()

	// Closure:
	// - A closure is a function that captures the environment in which it was created.
	// - A closure is a function that captures the variables in its lexical scope.

	// Example of a closure:
	// - func counter() func() int { count := 0; return func() int { count++; return count } }
	// Example calling a closure:
	// - c := counter(); fmt.Println(c()) // 1
	// - fmt.Println(c()) // 2
	// - fmt.Println(c()) // 3

	// Functions can be called from other functions:
	// - Functions in the same package can be called directly
	// - Functions in other packages must be called with the package name
	// The function name is "Hello" and it is in the "hello" package
	//var arrayFibonacci [50]int
	//complexdatatypes. .fibonacci(&arrayFibonacci)
}
