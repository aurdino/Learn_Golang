/*
this is multiple line comments

	Golang created by Google
	Created with aim
	-> the ease of programming of an interpreted, dynamically types language(eg python)
	-> with the efficiency and safety of a statically typed, complied lang(c++)
	-> it also aimed to be moder, with support for networked and multiple computing
	It is a complied programming lang -> sourcecode will be converted to machine understandable
	language

*/


package main

import "fmt" //form of format, used for -> input,output 

//this is a comment
func main() {   //function declaration
	name:= "Vishal"
    fmt.Println("Hi,",name) //print on terminal
	fmt.Println("What are you doing?")
}


/*
data types : defines a way the data stored, describe operations to be done on them, categories related values

String,
Number : Integer, float
Boolean : true, false

Array & Slices
[1,2,3,4]
["foo", "bar"]
[0.5,10.20,14.15]

Maps : collection of key-value pair
"x" -> 30
1 -> 100
"key" -> "value"

Memory allocation :
Integer  4 bytes(32-bit machine)
8 bytes (64-bit machine)
Boolean : 1 byte allocated

*/

/*
	Static vs Dynamically typed lang
	Static typed or strongly typed : 
		> compiler thows an error when data types are used incorrectly ,eg java, c++,c
		advantages : better performance, bugs can be caught by a compiler,better data integrity
	Dynamically typed or loosely type:
		>compiler does not enforce the type system, eg python,javascript
		no need to specify the data type during declaration of variables
		adv : faster to write code, less rigid, 
	go : 
	its fast,statically typed, compiled lang that feels like a dynamically typed, interpreted lang 

*/
