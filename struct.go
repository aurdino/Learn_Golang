package main

import (
	"fmt"
	"reflect"
)
/*
struct : a collection of fields ,groups data elements together
user data type
*/

type Vertex struct {
	X int
	Y int
}

type Student struct{
	name string
	rollNum string
	marks []int
	grades map[string]int
}

//pass struct to function
type Circle struct {
	x int 
	y int 
	radius float64
	area float64
}

func calArea(c *Circle) {
	var area float64
	const PI = 3.14
	area =  PI* c.radius * c.radius
	c.area = area
}

type s1 struct {
	x int
}

type s2 struct {
	x int
}


func main() {

	/*st1 := s1{x: 5}
	st2 := s2{x: 5}
	if st1 == st2 {
		fmt.Println("Matched")
	}*/

	//c := Circle(x: 5, y: 4, radius: 5, area: 0)

	//calArea(&c)
	//fmt.Println("c : ", c)

	var s Student
	//another way to initialize
	s1 := new(Student)
	fmt.Printf("%v %v\n", s,s1)
	//initialization 
	st := Student {
		name : "Vishal",
		rollNum : "BE1029514" ,
	}
	fmt.Printf("%+v %s: ", st,st.name)
	fmt.Println(Vertex{1, 2}, "\n", reflect.TypeOf(Vertex{2,3}))
	fmt.Printf("%T", Vertex{1,2})
}

