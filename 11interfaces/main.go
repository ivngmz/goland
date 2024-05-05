package _1interfaces

/*
Interfaces:
- An interface type is defined as a set of method signatures.
- A value of interface type can hold any value that implements those methods.
- A type implements an interface by implementing its methods. There is no explicit declaration of intent.
- Interfaces are implemented implicitly.
- A type implements an interface by implementing the methods defined in the interface.
- There is no explicit declaration of intent, no "implements" keyword.
- Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any
*/
type scheduler interface {
	getDays() int
	getHours() int
	getExtra() int
}

type employee struct {
	days  int
	hours int
	extra int
}

func (e employee) getDays() int {
	return e.days
}

func (e employee) getHours() int {
	return e.hours
}

func (e employee) getExtra() int {
	return e.extra
}

func main() {
	/* Interfaces:
	- An interface type is defined as a set of method signatures.
	- A value of interface type can hold any value that implements those methods.
	- A*/
	e := employee{5, 40, 10}
	var s scheduler = e
	println(s.getDays(), s.getHours(), s.getExtra())
}
