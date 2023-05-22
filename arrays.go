package main

import "fmt"

/**
	Array : collection of similar data elements stored at contiguous memory location
	fixed length after declaration we cannot change length 
	[1,2,3,4]
	200 204 208 212 memory locaiton
	["Ravi", "Rahul", "Anshu"]

	length
	capacity
*/
//sum of array
func sumArray(ar [10]int) int {
	
	sm := 0
	
	for _, num := range ar {
        sm += num
    }
	return sm
}

func main() {
	ar :=  [10]int{1,2,3,4,5}
	fmt.Println("sm: ",sumArray(ar))
	var grades [10]int
	fmt.Println(grades)

	var students [3]string = [3]string{"Ravi", "Rahul","Kamlesh"} //array initialization
	//students := [3]string{"12","34","56"}

	//students := [...]string{"12","132", "49u3"} //array with unspecified size 
	devices := [...]string{"Laptop", "Mobile","TV"}

	fmt.Println(devices)
	fmt.Println("length of array : ", len(devices))
	for i:= 0; i < 3; i++ {
		fmt.Println(devices[i])
	}
	for index, ele := range devices {
		fmt.Println(index, "=>", ele)
	}

	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

	/*for i:= 0; i <= 10; i++ {
		var x int
		fmt.Scanf("%d", &x)
		grades[i] = x
	}
	for i:= 0; i < 10; i++ {
		fmt.Println(grades[i])
	} */

	//multidimensional array
	arr := [3][2]int{{2,4}, {4,16}, {8,64}}
	fmt.Println(arr)
	//iterating 2-d array
	for i:= 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(" ", arr[i][j])
		}
		fmt.Print("\n")
	}
}