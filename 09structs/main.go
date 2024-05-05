package _9structs

import "fmt"

type worker struct {
	name     string
	surname  string
	age      int
	category string
	salary   int
}

func printWorker(w worker) {
	println("\n ##### Worker details with function:")
	fmt.Printf("Name: %s\n", w.name)
	fmt.Printf("Surname: %s\n", w.surname)
	fmt.Printf("Age: %d\n", w.age)
	fmt.Printf("Category: %s\n", w.category)
	fmt.Printf("Salary: %d\n", w.salary)
}
func printWorker2(w *worker) {
	println("\n ##### Worker details with function:")
	println("Salary increased inside function to 5000")
	w.salary = 5000
	fmt.Printf("Name: %s\n", w.name)
	fmt.Printf("Surname: %s\n", w.surname)
	fmt.Printf("Age: %d\n", w.age)
	fmt.Printf("Category: %s\n", w.category)
	fmt.Printf("Salary: %d\n", w.salary)
}

func main() {
	/*
		Structs:
			- Structs are used defined data types.
			- A structure groups together multiple data elements
			- Allows to reference series of grouped values through a unique variable name.
		Using fmt-Printf could be printed fields of Struct using %+v
	*/
	var worker1 worker
	fmt.Printf("%+v\n", worker1)
	worker2 := new(worker)
	fmt.Printf("%+v\n", worker2)
	//worker1 = new("Jim", "Koot", 44, "official", 2342)
	worker3 := worker{"Ataulfo", "Caderas", 33, "Electricist", 3122}
	fmt.Printf("%+v\n", worker3)
	worker4 := worker{
		name:     "Jim",
		surname:  "Koot",
		age:      44,
		category: "official",
		salary:   2342,
	}
	fmt.Printf("%+v\n", worker4)

	// Accessing fields of a struct
	fmt.Printf("Name of worker4 is: %s\n", worker4.name)
	fmt.Printf("Surname of worker4 is: %s\n", worker4.surname)

	// Modify fields of a struct
	fmt.Printf("Salary of worker4 is: %d\n", worker4.salary)
	worker4.salary = 3000
	fmt.Printf("Salary of worker4 has been changed to: %d\n", worker4.salary)

	// Passing struct to a function
	println("\n ##### Passing struct to a function by value")
	printWorker(worker4)

	// Passing struct to a function by reference
	println("\n ##### Passing struct to a function by reference")
	printWorker2(&worker4)
	println("\n ##### Worker salary details after function:\n")
	fmt.Printf("Salary of worker4 has been changed to: %d\n", worker4.salary)

	// Comparing structs
	println("\n ##### Comparing structs")
	worker5 := worker{
		name:     "Jim",
		surname:  "Koot",
		age:      44,
		category: "official",
		salary:   2342,
	}
	worker6 := worker{
		name:     "Jim",
		surname:  "Koot",
		age:      44,
		category: "official",
		salary:   2342,
	}
	if worker5 == worker6 {
		println("worker5 and worker6 are the same")
	} else {
		println("worker5 and worker6 are different")
	}
	if worker5 == worker4 {
		println("worker5 and worker4 are the same")
	} else {
		println("worker5 and worker4 are different")
	}
}
