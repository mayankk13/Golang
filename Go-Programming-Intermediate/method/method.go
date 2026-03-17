package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

type MyInt int

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

// method on user defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyOnt Type"
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Println("Area of rectangle is: ", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle is: ", area)

	num := MyInt(-5)
	fmt.Println(num.IsPositive())

	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{
		length: 10,
		width:  5,
	}}

	// as we are embeddimg a struct inside another struct, the method associated with thw embedded
	// struct will promoted to the outer struct directly, so we can can access area directly woth a dot.
	fmt.Println(s.Area())
}
