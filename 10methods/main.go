package _0methods

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	println(v.X, v.Y)
	return 0
}

type Cube struct {
	Vertex
	Z      float64
	volume float64
}

func (c *Cube) setVolume() {
	c.volume = c.X * c.Y * c.Z
}

func main() {
	/* Methods:
	- A method is a function with a special receiver argument.
	- The receiver appears in its own argument list between the func keyword and the method name.
	- In this example, the Abs method has a receiver of type Vertex named v.
	- You can only declare a method with a receiver whose type is defined in the same package as the method.
	- You cannot declare a method with a receiver whose type is defined in another package
		(which includes the built-in types such as int).
	- There are two reasons for this restriction:
		1. The method is defined for a specific named type, so it doesn't make sense to define a method on a built-in
			type.
		2. The method is defined in the same package as the named type, so it wouldn't be able to access the receiver
			if the receiver were defined in another package.
	- You can declare a method with a pointer receiver, but you cannot declare a method with a pointer receiver whose
		type is defined in another package.
	- You can declare a method with a value receiver whose type is defined in another package.
	*/
	v := Vertex{3, 4}
	v.Abs()
	fmt.Printf("%+v\n", v)
	c := Cube{v, 5, 0}
	fmt.Printf("%+v\n", c)
	c.setVolume()
	fmt.Printf("%+v\n", c)

}
