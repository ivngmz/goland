package _4operators

import "fmt"

func main() {
	// Arithmetic Operators
	fmt.Println("Arithmetic Operators:")
	a := 10
	b := 3
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Comparison Operators
	fmt.Println("\nComparison Operators:")
	fmt.Println("a == b =", a == b)
	fmt.Println("a != b =", a != b)
	fmt.Println("a < b =", a < b)
	fmt.Println("a > b =", a > b)
	fmt.Println("a <= b =", a <= b)
	fmt.Println("a >= b =", a >= b)

	// Logical Operators
	fmt.Println("\nLogical Operators:")
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	// Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	fmt.Println("a & b =", a&b)   // bitwise AND
	fmt.Println("a | b =", a|b)   // bitwise OR
	fmt.Println("a ^ b =", a^b)   // bitwise XOR
	fmt.Println("a << b =", a<<b) // left shift
	fmt.Println("a >> b =", a>>b) // right shift
}
