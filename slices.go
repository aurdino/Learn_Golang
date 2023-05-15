package main
import (
	"fmt"
	"reflect"
)

/*
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the 
elements of an array. In practice, slices are much more common than arrays.

->continuos segment of an underlying array
->variable typed(elements can be added or removed)
->more flexible

slices has 3 elements : pointer, length :len() : no of elements in slices, capacity:cap() 
pointer points to the first element 

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4]
*/

func main () {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]   //4 is not included ie high is not included
	fmt.Println(s)
	slice := []int{10,20,40}   //primes[:], complete array will be sliced 
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	p := primes[:]
	fmt.Print(" ", p)
	s1 := primes[:4]
	fmt.Print(" ", s1)
	fmt.Println("")
	//another way to declare slice
	s2 := make([]int,5,10) //make([]<data_type>, length, capacity)
	fmt.Println(s2)
	fmt.Println("len :", len(s2))
	fmt.Println("cap : ", cap(s2))

	arr := [10]int{10,20,30,40,50,60,70,80,90,100}
	/*in array length = capacity but in slice lenght != capacity, 
	capacity starts with 1 element of the slice 
	if we change an element in slice then this value also change in original array from which slice created
	-> slice has its own index
	*/

	s4 := arr[1:8]
	fmt.Println("capacity arr : ", cap(arr))
	fmt.Println("cap of slice : ", cap(s4))

	ar := [5]int{1,2,3,4,5}
	s5 := ar[:3]
	fmt.Println("before modification : ")
	fmt.Println("ar : ", ar)
	fmt.Println("s5 : ", s5)
	s5[0] = 10
	fmt.Println("After modification : ")
	fmt.Println("ar : ", ar)
	fmt.Println("s5 : ", s5)

	/*
		slice appending
		slice := append(slicename, elements)
		slice := append(slice, 10,20,40)
	*/

	fmt.Println("before slice appending : ")
	fmt.Println("s5 : ", s5)
	fmt.Println("s5 len: ", len(s5))
	fmt.Println("s5 cap: ", cap(s5))
	s6 := append(s5,10,20,30)
	fmt.Println("after slice appending : ")
	fmt.Println("s6 : ", s6)
	fmt.Println("s6 len: ", len(s6))
	fmt.Println("s6 cap: ", cap(s6))

	//appending to a slice 
	//slice = append(slice, anotherslice...) , delete, copy 

	a1 := [5]int{1,2,3,4,5}
	sl1 := a1[:3]
	a2 := [5]int{6,7,8,9,10}
	sl2 := a2[:3]
	new_slice := append(sl1, sl2...)
	fmt.Println(new_slice)
	
	//copy 
	src_slice := []int{10,20,40,30,50}
	dest_slice := make([]int,3)
	num := copy(dest_slice, src_slice)
	fmt.Println("dest slice : ", dest_slice)
	fmt.Println("copied elements : ", num)

	for _, value := range src_slice {
		fmt.Println(value)
	}

	sl3 := []int{1,2,3,4,5,6}
	fmt.Println("typef : ", reflect.TypeOf(sl3))
	fmt.Println("sl3 : ", sl3)
	sl4 := append(sl3, 7,8,9)
	fmt.Println("sl4 : ", sl4)

	sl5 := []int{}
	sl6 := append(sl5, 1,2,3,4)
	fmt.Println("sl6 : ", sl6)

}
