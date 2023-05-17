package main

import (
	"fmt"
)
/*
passing by value : function is called by directly passing the value of the variable as an argument
->parameter is copied to into another location of your memory
->so when accessing or modifying the variable within your function, only the copy is accessed or modified 
and original value is never modified
-> int,string,bool,array are passed by value

call by reference : go supports pointers,allowing you to pass references to values within your prog
-> in this, the address of the variable is passed into the function call as the actual parameter
->All the operations in the function are performed on the values stored at the address of the actual parameter

*/

//call by value
func modify (a int ) {
	a += 100
}

//call by reference

//** slices and maps are passed by reference by default

func modify3(arr [] int) {
	arr[0] = 100
}

func modifiedRef(a *int) {
	*a += 100
}
func modifyMap(m map[string]int) {
	m["One"] = 2
}

func main() {
	a := 10 
	fmt.Println("a : ", a)
	modify(a) //value of a will not modify 
	fmt.Println("modified a : ", a)

	modifiedRef(&a)
	fmt.Println("modified ref a : ", a) // changed to 110

	arr := []int {10,20,30}
	fmt.Println("initial slice : ", arr)
	modify3(arr)
	fmt.Println("modified slices : ", arr) 

	//map is also pass by reference by default

	mp := make(map[string]int)
	mp["One"] = 1
	mp["Two"] = 2
	mp["Three"] = 3
	fmt.Println("mp :", mp)
	modifyMap(mp)
	fmt.Println("mp modified :", mp)
}
