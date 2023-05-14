package main

import "fmt"

func main() {
	var s string = "Happy"
	if s == "happy" {
		fmt.Println("I am happy!")
	} else { //need to right else in the same line where if closing 
		fmt.Println("Not happy")
	}
	
	fruit := "grapes"

	if fruit == "apple"{
		fmt.Println("Apple")
	} else if fruit == "orange" {
		fmt.Println("orange")
	} else {
		fmt.Println("not appetite")
	}

	//switch case

	var i int = 100

	switch i	{
	case 10 :
		fmt.Println("i is 10")
	case 100,200:
		fmt.Println("i is either 100 or 200")
	default:
		fmt.Println("is neither 0, 100 or 200")
	}

	var j int = 800

	switch j	{
	case 10 :
		fmt.Println("j is 10")
	case 100,200:
		fmt.Println("j is either 100 or 200")
	default:
		fmt.Println("j is neither 0, 100 or 200")
	}

	//fallthrough : used to force the execution flow to fall through the successive case block
	//similar to continue in c++, if condition matched, after that all below conditions will be executed


	var k int = 10

	switch k {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("default")

	}

	var a,b int = 10,20
	switch {
	case a+b == 30:
		fmt.Println("equal to 30")
	case a+b <= 30:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}

	/*
		Looping

	*/
	for i := 1; i <= 5; i++ {
		fmt.Println(i*i)
	}
	//or 
	l := 1
	for l <= 5 {
		fmt.Println(l*l)
		l += 1
	}
	//infinite loop 
	/*sm := 0
	for {
		sm++
	}
	fmt.Println(sm)
	*/

	//break : ends loop immediately when it is encountered

	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("\n")
	//continue : skips the current iteration of loop and continues with the next iteration

	for i:= 1 ; i <= 5; i++ {
		if i == 3 {
			continue //skips 3 
		}
		fmt.Println(i)
	}
}