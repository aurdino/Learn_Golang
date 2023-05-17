package main

import (
	"fmt"
)
/*
 defer statement : delays the execution of a function until its surrounding functions executed
 deferred call's arguments are evaluated immediately, but the function call is not executed until the 
 surrounding function returns
*/

func printName(name string){
	fmt.Println("Name :", name)
}

func printRollno(rolnum string){
	fmt.Println("Roll number : ", rolnum)
}

func printAddress(addr string){
	fmt.Println("Address : ", addr)
}

func main () {
	printName("Vishal") //sequence of execution of functions : 1,3,2
	defer printRollno("BE/10295/14")
	printAddress("Siwan, Bihar")
}

/*
func addHundred(x int) int {
	return x + 100
}
func partialSum(x ...int) func() int {
	sum := 0
	for _, value := range x {
			sum += value
	}
	return func() int {
			return addHundred(sum)
	}
}
func main() {
	partial := partialSum(1, 2, 3)
	partial() //no output 



prog2 
func addHundred(x int) int {
        return x + 100
}
func partialSum(add100 func(x int) int, x ...int) int {
        sum := 0
        for _, value := range x {
                sum += value
        }
        return add100(sum)

}
func main() {
        partial := partialSum(addHundred, 1, 2, 3)
        fmt.Println(partial) //106
}
prog3 : 


func addHundred(x int) {
        fmt.Println(x + 100)
}
func partialSum(add100 func(x int), x ...int) int {
        sum := 0
        for _, value := range x {
                sum += value
        }
        add100(sum)
        return 0
}
func main() {
        partial := partialSum(addHundred, 1, 2, 3)
        fmt.Println(partial) //output : 106 0
}

}*/