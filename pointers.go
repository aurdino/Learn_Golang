package main 

import (
	"fmt"
)
/*
	1.pointer : is a variable that holds the memory address of another variable
	->they point where the memory is allocated and provide ways to find or even change the 
	value located at the memory location

	2.address and dereference operator

	& operator : the address of a variable can be obtained by preceding the name of a variable
	with an &(ampersand) sign, called address-of-operator.

	* operator (dereference) -> when placed before an address, it returns the value at that address





*/
func main() {
	s := "hello"
	var b *string = &s
	fmt.Println("b : ", b)
	var a = &s
	fmt.Println("a : ", a)

	i := 10 
	fmt.Printf("%T %v\n", &i, &i)
	fmt.Printf("%T %v\n", *(&i), *(&i))

	var sn *string //pointer declaration
	var j *int //pointer declaration
	fmt.Println(sn, "\n",j) //nil denotes zero in case of int, empty string("") in case of string
	//initialization of pointer : var <ptr_name> *<data_type> = <&variable_name> or var <ptr_name> = &<var_name>

	/*k := 10
	var ptr_k *int = &k //stores address of k
	//or
   var s string = "hello"
   var ptr_s = &s
   fmt.Println(*s, *ptr_k,*ptr_s) */

}