package _5controlflow

import "fmt"

func main() {
	// if statement
	num := 2
	if num%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	// switch statement
	switch num {
	case 1:
		fmt.Println("The number is 1.")
	case 2:
		fmt.Println("The number is 2.")
	case 3:
		fmt.Println("The number is 3.")
	default:
		fmt.Println("The number is not 1, 2, or 3.")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop with range
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	/* break statement
	- The break statement is used to terminate the loop immediately.
	*/
	for k := 0; k < 5; k++ {
		if k == 3 {
			break
		}
		fmt.Println(k)
	}

	// continue statement
	for l := 0; l < 5; l++ {
		if l == 3 {
			continue
		}
		fmt.Println(l)
	}

	// nested loop
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			fmt.Println(m, n)
		}
	}

	/* labeled loop
	- A labeled loop is a loop that has a label.
	- The label is used to identify the loop.
	- The break and continue statements can be used with the label to break or continue the loop.
	- The label is placed before the loop.
	- The label is followed by a colon.
	*/

outerLoop:
	for o := 0; o < 3; o++ {
		for p := 0; p < 3; p++ {
			fmt.Println(o, p)
			if o == 1 && p == 1 {
				break outerLoop
			}
		}
	}

	/*  defer statement
	- The defer statement is used to delay the execution of a statement until the surrounding function returns.
	*/
	defer fmt.Println("This will be executed last.")
	fmt.Println("This will be executed first.")

	/* panic statement
	- The panic statement is used to terminate the program immediately.
	- The panic statement is used to handle unexpected errors.
	*/
	// panic("This is a panic.")

}
