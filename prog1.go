package main

import "fmt"

func main() {
	var firstName string = "Vishal"
	var lastName string = "Rajak"
	fmt.Println("Your complete name is : ", firstName + " " +  lastName)
	var a int = 15
	var b int = 10
	fmt.Print(a)
	fmt.Print(b) //print value in single line concatenately, 1510
	fmt.Println("one" + "plus") //prints value in new line
	fmt.Println("one",  1)
	fmt.Println(7.0/3.0) //float value
	fmt.Println(true && false) // false
	fmt.Println(true || false) //true 
	fmt.Println(true && true) // true 
	fmt.Println("value of a : ", a, "\n", "Value of b : ", b, "\n","sum of a and b : ", a+b)
	fmt.Println(a/b)

	fmt.Println("Testing to push this file on github")

	//Printf - format specifier
	//%v formats the value in a default format
	//%d used for integers
	//%c : character 
	//%s : string
	//%t : true or false
	//%f : floating numbers
	//%.2f : floating numbers upto 2 decimal places

	var s string = "Golang"
	fmt.Printf("Welcome to Golang, %v", s)
	var val int = 100
	fmt.Printf("val : %d",val)
	var age int
	age = 25 
	fmt.Println("age : ", age)
	var f,t = "foo", "too"
	fmt.Println(f,t)

	var (
		i int = 10
		sr string = "this is a string"
	)
	fmt.Println(i, sr)

	//short variable declaration 
	fn := "This is short string"
	fmt.Println(fn)

	rate := 10 
	fmt.Print("rate : ", rate)
	//rate := "Vishal"

	//fmt.Print("ratestrig", rate) //error
	rate := 40  //error
	fmt.Println("rate3 : ", rate)
}

/*
data types : 
unsigned integers : contain only positive numbers including zero
uint8 unsigned integer //8bits or 1 byte
uint16 : 2bytes or 16 bits 
uint32 : 4 bytes 
uint64 : 8 bytes 

int8 : signed integers -> both negative and positives numbers
int16 : 2 bytes 
int32 : 4 bytes 
int64 : 8 bytes
int : 4 bytes for 32-bit machines, 8 bytes for 64-bit

Float : floating point number
float32: 32 bits or 4 bytes 
float64 : 64 bits or 8 bytes
0.10
68.10

string:  sequence of characters numbers, letters, symbols
"abc", "0102a"
16 bytes of memory occupied by string in go

Boolean : 
bool : only true or false //1 byte 
*/

/*
Variables : used to assign a value 
*/