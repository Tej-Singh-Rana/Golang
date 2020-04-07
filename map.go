// map function :- A map is an unordered collection of key-value pairs. 
// Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key.
// The type of keys and type of values must be of the same types, different types of keys and values in the same maps are not allowed.
// In the maps, a key must be unique.
// var variable-name mape[key-type]value-type
package main

import "fmt"

func main(){
	// initializing map
	var map_1 = map[string]string{"car":"Bmw","model":"UG29VGL","year":"2020","color":"blue",}

	fmt.Println("Element is : ", map_1)
	fmt.Println("Car is : ", map_1["car"])


}