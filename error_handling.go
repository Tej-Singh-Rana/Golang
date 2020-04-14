// error handing
//

package main

import (
			"fmt"
			"os"
)

func main(){
			// replace with your file location.
		 // _(underscore) ignoring value, that's not important at this moment.
			_, err := os.Open("C:\\iptables_installation_configure.txt")
			// if error value is not nil then print error details.
			if err != nil{
				fmt.Println("Error returned was :",err)
			}
			// file available so returned nil value.
			fmt.Println("Error returned was :",err)
}