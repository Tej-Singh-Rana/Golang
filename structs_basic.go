// structs --> Custom data types.
/*
  structs have a specific type and used to define a new type. It's field can be any type.
  structs are basically a collection of named fields, each field of a specific type. 
  each field in a struct can be of a different type(int,string,float) to the other fields in the struct.

*/

package main

import "fmt"

func main(){
			// declare and initialize, keyword, type-name, define that new type-name is going to be struct.
			// custom types
			// keyword type type-name struct{ name fields / metadata }
      type bookDetail struct {
        
        // bookName   string
        // language   string
				bookName, language   string
				price			float64
			}
			// var book bookDetail
			// "new" keyword is give us to a pointer
			// 		book := new(bookDetail)
			// 		*bookDetail
			// This allocates memory for all the fields, sets each of them to their zero value and returns a pointer.

			// declaring a new variables
					book := bookDetail {
					bookName: "Object-oriented",
					language: "English",
					price: 245.50,
			}
			fmt.Println(book)
			// period operator to define individual fields
			// accessing the fields using the dot(.) operator
			fmt.Println("\nName of the book is :",book.bookName)
			fmt.Println("\nLanguage is book written :",book.language)
}