package main

import "fmt"

/*
	inner blocks can access variables declared within outer blocks 
	while outer blocks cannot access variables declared within inner blocks

	{
		//outer block 

		{
			//inner block
		}

	}
	local variables : declared inside a function or a block can be access within that function or block
	and not accessible outside the block,also defined inside looping or conditional statements
	global variables : declared outside the funciton or block
	can be access throughout the lifetime of a program
	
*/

var name string = "Vishal" //global variable
var age int = 26 //age := 26 will not work
func main() {

	city := "New Delhi" //local variable

	{
		country := "India"
		fmt.Println(country)
		fmt.Println(city)
		fmt.Println(name)
		fmt.Println(age)
	}
	//fmt.Println(country) //here accessing inner block code, will give error

	//zero values -> by default contains below values if we do not define any value to them
	/*
		bool : false, int : 0, float : 0.0, string = ""
	*/

	var i int 
	var f float64
	var b bool 
	//var s string
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(b)
}