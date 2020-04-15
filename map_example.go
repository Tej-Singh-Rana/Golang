// all keys in a given map have got to be unique.
// we can't have two items in a map with the same key.
// map[key-type]value-type
// map is unordered list
// possible scenario insert, update, delete with map items.


package main

import "fmt"

func main(){
				// declare map func
        countDown := make(map[string]int)

        countDown["one"] = 120
        countDown["two"] = 150

        fmt.Println(countDown)

        countUp := map[int]string{1: "Florida", 2:"Connecticut", 3:"Ohio", 4:"California"}

        for key, value := range countUp {
							// unordered list pairs
              fmt.Println("\nKey is:",key,"Value is:",value)
				}
				//--------------insert/update/delete----------
				numCount := map[int]string{1:"good", 2:"flexible", 3:"unordered", 4:"list"}

				numCount[3] = "ping"    // updated data value
				numCount[5] = "unordered"   // insert new data value
				fmt.Println("\n",numCount)
				delete(numCount,4)      // delete 
				fmt.Println(numCount)
}