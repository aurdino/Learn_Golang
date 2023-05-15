package main

import (
	"fmt"
)

/*
	Map : 
		unordered collection of key/value pair, also called associative array,dictionary
		implemented by hash table
		provide efficient add,get and delete operations

*/
func main () {
	var codes map[string]string //var <map_name>map[key_data_type]<value_data_type> nil map 
	//codes["en"] = "English"
	fmt.Println("codes :", codes)  //we cannot add value to a nil map 

	//<map_name> := map[key_data_type]<value_data_type>{key-value-pairs}

	mp := map[string]string{"en" : "English", "hn" : "Hindi", "fr" : "French"}

	fmt.Println("map : ", mp)
	fmt.Println(mp["en"])
	//we also create map using make() , <map_name> :=make(map[key_data_type]<value_Data_type>,<initial_capacity>)
	
	code := make(map[string]int)
	code["one"] = 1
	code["two"] = 2
	code["three"] = 3
	fmt.Println("code:",code)
	value,found := code["one"]
	fmt.Println(value,found)
	//delete(code,"one")

	for key,value := range code {
		fmt.Println(key, "=>", value)
	}

	//truncate a map : clearing all elements from the map

	for key,value := code {
		delete(code, key)
	}
	fmt.Println(code)//empty map
	//re-initialize map with empty map 
	code := make(map[string]int) //map is truncated

	/*
	pro1
	var ascii_codes map[string]int
    ascii_codes["A"] = 65
    fmt.Println(ascii_codes) //error 

	prog2 
	var ascii_codes map[string]int
    ascii_codes["A"] = 65
	fmt.Println(ascii_codes) //error 
	prog3 :
	ascii_codes := map[string]int{}
        ascii_codes["A"] = 65
        _, found := ascii_codes["B"]
        if found {
                fmt.Println("key B was not found")
        } //no output 

		Prog 4 :
		ascii_codes := make(map[string]int)
        ascii_codes["A"] = 65
        ascii_codes["F"] = 70
        ascii_codes["K"] = 75
        fmt.Println(ascii_codes)

        ascii_codes = make(map[string]int)
        ascii_codes["U"] = 85
        fmt.Println(ascii_codes)
	*/
}