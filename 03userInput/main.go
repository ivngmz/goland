package _3userInput

import "fmt"

func main() {
	var num1, num2, result int
	num1, _ = fmt.Scanf("Enter a number: ")
	num2, _ = fmt.Scanf("Enter another number: ")
	result = num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, result)
}
