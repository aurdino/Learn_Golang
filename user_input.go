package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string 
	var age int
	fmt.Print("Enter your name and age")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Print("Your name is : ", name, age)
	var is_muggle bool 
	fmt.Scanf("%t", &is_muggle)
	fmt.Print(is_muggle)
	/*
	 Scanf returns  : count -> the number of arguments that the function writes to 
	 err : any error thrown during the execution of the function

	*/

	var a string 
	var b int

	fmt.Print("Enter a string and a number")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count : ", count)
	fmt.Println("error : ", err)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	/*
		Variable Types: %T format specifier
		reflect.TypeOf
	*/
	fmt.Printf("type of %v is %T", a, a)
	fmt.Printf("Type of %v is %T", b,b)
	fmt.Printf("Type : %v \n", reflect.TypeOf(10))
}
