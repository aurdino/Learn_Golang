package main

import (
	"fmt"
)

/*
	function parameters : are the names listed in the function's definition
	while function arguments are the real values passed into the function
*/
//add two numbers 
func addTwoNumbers(a,b int) int {
	return a+b
}

//check a number is prime or not 

func checkPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num % i == 0 {
			return false 
		}
	}
	return true
}

//sieve of Eratosthenes 

func sieve(num int) []int {

	check := make([]bool, num+1)
	for i := 0; i <= num; i++ {
		check[i] = true
	}

	for i:= 2; i*i <= num; i++ {
		if check[i] == true {
			for j:= i*2; j <= num; j += i {
				check[j] = false 
			}
		}
	}
	var primes []int
	for i := 2; i <= num; i++ {
		if(check[i]) {
			primes = append(primes,i)
		}
	}
	return primes
}

//a function can return multiple value 
func operation (a int, b int) (int ,int) {
	 sum := a+b
	 diff := a-b
	 return  sum,diff// or return sum,diff 
}

//variadic function : accept variable number of arguments
func addNum(numbers ...int) int {
	sm := 0
	for _,val := range numbers {
		sm += val
	}
	return sm
}

//blank identifier

func f() (int,int) {
	return 10,53
}
func main () {
	v,i  := f()
	fmt.Println(v,i)
	//v,_ := f(),fmt.Println(v) 
	fmt.Println(addNum(1,2,3,4,5))
	sm, diff  := operation(10,5)
	fmt.Println(sm, diff)
	var num int
	fmt.Scanf("%d", &num)
	if checkPrime(num) {
		fmt.Println(num , " is a prime number")
	} else {
		fmt.Println(num , "is not a prime number")
	}
	sum := addTwoNumbers(10,20)
	fmt.Println("sum : ", sum)
	fmt.Println(sieve(20))
}
