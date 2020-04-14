// env variables in Go

package main

import ( 
		"fmt"
		"os"
)

func main(){
		// env variable values in key-value pairs
			// fmt.Println("env variable values :",os.Environ())
			// index , value
			for vale, env := range os.Environ(){
				fmt.Println(vale,env)
			}
			// get the user name
			member := os.Getenv("username")
			fmt.Println("Username of machine :",member)
			

		}