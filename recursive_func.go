package main

import (
	"fmt"
)
/*
	A function which calls itself is called a recursive function
	-> it keeps calling itself until a base condition matched
	->used to solved a problem where the solution is dependent on smaller instance of the same problem

*/

func fact(num int ) int {
	if num == 1 {
		return 1
	}
	return num * fact(num-1)
}

/*
	Anonymous function : declared without any name identifier
	they can accept input and return outputs, just as standard functions do.
	used for containing functionality that need not be named and possibly for short-term use.

*/

func main() {
	var num int
	fmt.Println("Enter a number to find its factorial : ")
	fmt.Scanf("%d", &num)
	fa := fact(num)
	fmt.Println("factorial of ", num, " is : ", fa)

	//anonymous function 
	x := func(l int, b int) int { //here x data type is function 
		return l*b
	}
	prod := x(10,5)
	fmt.Println(prod)
	//2nd way

	y := func(l int, b int) int {  //here y data type is int
		return l*b 
	} (5,19)
	fmt.Println("y : ", y)
}