package main

import "fmt"


/*
	An interface specifies a method set and is a powerful way to introduce modularity in Go.
	interfaces is like a blueprint for a method set 
	They describe all the methods of a method set by providing the function signature for each method
	They specify a set of method but do not implement them
	->interfaces are implemented implicitly
	type <interface_name> interface {
		//methos signature
	}
*/
/*

	type Student interface {
        getPercentage() int
        getName() string
}

type Undergrad struct {
        name   string
        grades []int
}

type Postgrad struct {
        name   string
        grades []int
}

func (p Postgrad) getPercentage() int {
        sum := 0
        for _, v := range p.grades {
                sum += v
        }
        return ((sum * 100) / (len(p.grades) * 200))
}
func (p Postgrad) getName() string {
        return p.name
}

func (u Undergrad) getPercentage() int {
        sum := 0
        for _, v := range u.grades {
                sum += v
        }
        return sum / len(u.grades)
}
func (u Undergrad) getName() string {
        return u.name
}

func printData(s Student) {
        fmt.Println(s.getName())
        fmt.Println(s.getPercentage())
}

func main() {
        u := Undergrad{"Ross", []int{90, 75, 80}}
        p := Postgrad{"Joe", []int{150, 190, 185}}
        printData(u)
        printData(p)
}

*/

type shape interface {
	area() float64
	perimeter() float64
}
type square struct {
	side float64 
}
func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4*s.side
}

type rect struct {
	length,breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())


}
func main() {
	r := rect {length: 3, breadth: 4}
	c := square{side: 5}
	printData(r)
	printData(c)
}