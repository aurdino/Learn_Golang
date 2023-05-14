package main 

import "fmt"

/*
		types of operators : 
		comparison : ==,!=,<,<=,>,>=, allows comparison for same data types 
		arithmetic : +, -, *, /,%,++,--
		assignment operator : =,+=,-=,*=,/=,%=
		bitwise operator : &, | ,<<, >>,^(xor)
		logical operators : !, &&, || 
	*/

func main() {
	city1 := "Kolkata"
	city2 := "Calcutta"
	fmt.Println(city1 == city2) //false 
	fmt.Println(city1 != city2) // true
	var a,b int = 5,10
	fmt.Println(a < b, "\n", a != b, "\n", b > a)

	var x int = 10
	x++;//x+1
	fmt.Println("x : ", x)
	var y int = 5;
	x += y
	fmt.Println("x : ", x)
	y-- // y-1
	fmt.Println("y : ", y)
	x *= y;
	fmt.Println("x*y : ", x)

	//bitwise operators
	/*
		&(AND), |(OR) , ^(XOR), >>(right shift), <<(left shift)
		& : takes two numbers as operands and does AND on every bit of two numbers
		12 : 01100
		25:  11001
		12&25 : 01000 -> 8

		| : takes two numbers as operands and does OR on every bit of two numbers
		12 : 01100
		25:  11001
		12&25 : 11101 -> 29

		^ : takes two numbers as operands and does XOR on every bit of two numbers
		the result of XOR is 1 if the two bits are opposite

		12 : 01100
		25:  11001
		12&25 : 10101 -> 21

		<<(left shift) : shifts all bits towards left by a certain number of specified bits
		the bit positions that have been vacated by the left shift operator are filled with 0
		the result of XOR is 1 if the two bits are opposite

		212 : 11010100
		212 << 1 -> 424
		11010100
	   110101000

	   10 -> 1010
	   10 << 1 -> 20
	   1010
	  10100 -> 20

	  >> (right shift) : shifts all bits towards right by a certain number of specified bits
	  excess bits shifted off to the right are discarded

	  212 -> 11010100
	  212 >> 2 -> 53
	  11010100
	  00110101 -> 53 

	  8 -> 1000
	  8 >> 2 -> 0010 -> 2  (x >> n = x/pow(2,n)))
	*/
	fmt.Println(12&25, 12 | 25, 12^25, 212 << 1, 10 << 1, 212 >> 2,8>>3)
}