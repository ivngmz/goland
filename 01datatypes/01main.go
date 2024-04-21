package golang

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
	// String
	var name = "Andromeda"
	name2 := "Milky Way"
	// Number
	var age int = 13            // Age of the universe in billions of years
	var height float64 = 100000 // Average galaxy diameter in light years
	// Boolean
	var isSpiral bool = true
	// Array
	var galaxies = [3]string{"Andromeda", "Milky Way", "Triangulum"}
	// Slice
	var colors = []string{"Red", "Green", "Blue"} // Colors of galaxies
	// Map
	var galaxy = map[string]string{
		"name":     "Andromeda",
		"age":      "10",  // Age in billions of years
		"diameter": "220", // Diameter in thousands of light years
		"isSpiral": "true",
	}

	// Print the data type
	fmt.Printf("Name: %T\n", name)
	fmt.Printf("Name2: %T\n", name2)
	fmt.Printf("Age: %T\n", age)
	fmt.Printf("Diameter: %T\n", height)
	fmt.Printf("Is Spiral: %T\n", isSpiral)
	fmt.Printf("Galaxies: %T\n", galaxies)
	fmt.Printf("Colors: %T\n", colors)
	fmt.Printf("Galaxy: %T\n", galaxy)

}
