package main

import "fmt"

type s1 struct {
	x int
}

type s2 struct {
	x int 
}

type Movie struct {
	name   string
	rating float32
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Student struct {
	name string
	course string 
	age int 
	address string 
}
var (
	s_1 = Student{"Vishal","CSE",26,"Siwan"}
	s_2 = Student{course:"CSE", age:26}
	s_3 = Student{}
)


func getMovie(s string, r float32) (m Movie) {
	m = Movie{
			name:   s,
			rating: r,
	}
	return
}

func increaseRating(m *Movie) {
	m.rating += 1.0
}


func main() {
	/*mov := getMovie("xyz", 2.0)
	increaseRating(mov)
	fmt.Printf("%+v", mov) //error
	*/
/*	st1 := s1{x: 5}
	st2 := s2{x: 6}

	if st1 == st2 { //cannot compare two struct
		fmt.Println("matched")
	}*/
	fmt.Println(s_1, s_2, s_3)
	fmt.Println(v1, p, v2, v3)

	s := s1{x: 5}
	s2 := s1{x: 6}
	s3 := s1{x: 5}
	if s != s2 {
		fmt.Println("s and s2 have different value")
	}
	if s == s3 {
		fmt.Println("s and s3 have same value")
	}
}