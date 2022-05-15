//
//

package main

import (
      "fmt"
      
)

func main(){
			// slice method
			tech := []string{
				"kubernetes",
				"docker",
				"jenkins",
				"ansible",
				"linux",
				"grafana",
				"prometheus",
				"rancher",
}
			// index value, data value := range var-name
		// for _, i := range tech{
			for index, value := range tech{

			fmt.Println(index,value)
		}
		

}