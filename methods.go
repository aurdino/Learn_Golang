package main 

import "fmt"

/*
	Method : a method is a function that has a defined receiver
	func <receiver> <method_name>(<paramters>) <return_param>{
	}
	func (c Circle) area() float64 {

	}
*/

type Circle struct {
	radius float64
	area float64
}

func (c *Circle) calArea() {
	c.area = 3.14 * c.radius * c.radius
}

type Crcl struct {
	radius float64
	area float64
}
func (c Crcl) calArea1() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}
	c.calArea()
	fmt.Printf("%+v", c)

	c1 := Crcl{radius: 5}
	c1.calArea1() //area is zero here,call by val
	fmt.Printf("%+v", c1) 
}