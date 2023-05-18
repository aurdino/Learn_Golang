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