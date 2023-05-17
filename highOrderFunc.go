package main

import (
	"fmt"
)
/*
High Order Function : a function that receives a function as an argument or returns a function as output
->why use -> 1.composition : creating smaller functions that takes care of certain piece of logic
			composing complex function by using different smaller functions
	2.reduces bugs : 
	3.code get easy to learn and understand

*/

func calArea(radius float64) float64 {
	return 3.14*radius*radius
}

func calPerimeter (radius float64) float64 {
	return 2*3.14*radius
}

func calDiameter (radius float64) float64 {
	return 2*radius
}

/*func printResult(radius float64, calFunction func (r float64) float64 ) {
	result := calFunction(radius)
	fmt.Println("Result : ", result)
}

func getFunction (query int) func (r float64) float64 {
	query_to_func := map[int]func(r float64) float64 {
		1: Area,
		2: Perimenter,
		3:Diameter,
	}
	return query_to_func[query]
}
*/

func sum(x,y int) int {
	return x+y
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x,y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main () {
	res := partialSum(20)
	fmt.Println(res(10))
	fmt.Println(squareSum(3)(4)(5))
	var query int
	var radius float64
	fmt.Print("Enter the radius of circle and a query \n 1-Area \n 2-Perimeter \n 3-diameter : ")
	fmt.Scanf("%f %d", &radius, &query)
	fmt.Printf("") 

	if query == 1 {
		fmt.Println("Area : ", calArea(radius))
	} else if query == 2 {
		fmt.Println("Perimenter : ", calPerimeter(radius))
	} else if query == 3 {
		fmt.Println("Diameter : ", calDiameter(radius))
	} else {
		fmt.Println("Invalid input")
	}
	//printResult(radius, getFunction(query))

}