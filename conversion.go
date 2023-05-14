package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main () {
	//converting int to float
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	var f1 float64 = 10.34
	var j int = int(f1)
	fmt.Printf("%d\n", j) //float to integer

	//conversion of string using strconv package
	//Itoa() -> converts int to string, return one value -string formed with the given integers 
	var s string  = strconv.Itoa(i)
	fmt.Printf("%q %v\n", s, reflect.TypeOf(s))

	//Atoi() : converts string to int , returns two values - the corresponding integer, error(if any)

	var name string = "200"
	
	k,err := strconv.Atoi(name)
	fmt.Printf("%v %T\n", k,k)
	fmt.Printf("%v %T\n", err,err)

}