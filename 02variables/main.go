package _2variables

import (
	"fmt"
	"strings"
)

func main() {

	var country string = "Indonesia"
	var city string = "Jakarta"
	var citizens int = 10000000
	var isCapital bool = true
	var latitude float32 = 6.2088

	fmt.Println("the city of ", city,
		"is located on the country of ", country,
		"\nDo you know it?")

	/*
		- PrintF is a function that is used to format the output.
		- Options:
			- %v: Value in a default format.
			- %s: Value as a string.
			- %d: Value as a decimal integer.
			- %f: Value as a floating-point number.
			- %2f: Value as a floating-point number with 2 decimal points.
			- %t: Value as a boolean.
			- %T: Type of the value.
			- %p: Pointer.
			- %c: Character.
			- %x: Hexadecimal.
			- %e: Scientific notation.
			- %q: Quoted string.
			- %b: Binary.
			- %o: Octal.
			- %U: Unicode.
			- %v: Value in a default format.
			- %q: Quoted string.
			- %x: Hexadecimal.
			- %X: Hexadecimal with capital letters.
			- %U: Unicode.
	*/
	fmt.Printf("The city of %v is located on the country of %s. \n"+
		"The total number of citizens is %d. The city is located at latitude %f\n"+
		". Is it the capital city? %t\n"+
		"\nDo you know it?\n",
		strings.ToTitle(city), strings.ToTitle(country), citizens, latitude, isCapital)

	/*
		- Variable scope: The scope of a variable is the region of the program
			where the variable is accessible.
			- Go has block-level scope.
			- A variable declared inside a block is only accessible within that block.
			- A variable declared outside a block is accessible throughout the program.
		- Local variables:
			- Local variables in go are declared inside a block.
			- Local variables are only accessible within that block.
		- Global variables:
			- Global variables in go are declared outside a block.
			- Global variables are accessible throughout the program.
			- are available during the entire lifetime of the program.
		- Zero value: The default value of a variable if no value is assigned.
			- bool: false
			- int: 0
			- float: 0.0
			- string: ""
		- Short declaration: A short declaration is used to declare and initialize a variable.
			- Syntax: name := value
		- Type inference: The ability of the compiler to determine the data type of a variable based on the value assigned to it.
		- Type conversion: The process of converting a value from one data type to another.
			- Type casting: The process of converting a value from one data type to another.
			- Syntax: type(value)
		- Constants: A constant is a variable whose value cannot be changed.
			- Constants are declared using the const keyword.
			- Constants must be initialized at the time of declaration.
			- Constants can be of any data type.
			- Constants are immutable.
			- Example: const pi = 3.14
		- Enumerated constants: Enumerated constants are a set of named constants.
			- Example:
				const (day1 = "Monday", day2 = "Tuesday", day3 = "Wednesday")
		- iota: The iota identifier is used to create a set of related constants.
			- The value of iota starts at 0 and increments by 1 for each subsequent constant.
			- Example: const (day1 = iota, day2, day3)
		- Multiple short declarations: Multiple short declarations can be done in a single line.
			- Multiple short declarations: Multiple short declarations can be done in a single line.
			- Syntax: name1, name2 := value1, value2
		- Multiple variable declarations: Multiple variables can be declared in a single line.
			- Syntax: var name1, name2 type = value1, value2
	*/
	cityNearest := "Bogor"
	{
		country := "Malaysia"
		fmt.Printf("The nearest city of Indonesia to %s is %v\n", country, cityNearest)
	}
}
