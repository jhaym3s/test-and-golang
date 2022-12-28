package geometry

import "math"

///A method is a function with a receiver. A method declaration
//binds an identifier, the method name,
//to a method, and associates the method with the receiver's base type.

type Rectangle struct{
	height float64
	width float64
}
type Circle struct{
	radius float64
}

type Square struct{
	side float64
}
 type Shape interface{
	Area() float64
 }

func (r Rectangle) Perimeter() float64 {
	return 2*(r.height+r.width)
}
 
func (c Circle) Area() float64{
	return math.Pi * (c.radius*c.radius)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (s Square) Area() float64  {
	return s.side * s.side
}