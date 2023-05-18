package main

import "fmt"

/*
	Method Sets: set of methods that are available to a data type
	a way to encapsulate functionality

*/

type Student struct {
	name string 
	grades []int
}

func (s *Student) displayName() {
	fmt.Println("Name: ", s.name)
}

func (s *Student) calPercentage() float64 {
	sm := 0
	for _, val := range s.grades {
		sm += val
	}
	return float64(sm*100) / float64(len(s.grades)*100)
}

func main() {
	s := Student{name: "Vishal", grades: []int{90,75,80}}
	s.displayName()
	fmt.Printf("%.2f", s.calPercentage())
}