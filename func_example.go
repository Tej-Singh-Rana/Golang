//
//

package main

import (
		"fmt"
		"strings"
)

var (
				name = "Virtual Box"
				version = "latest"

)

func main(){

				fmt.Println(excode(name,version))
}
func excode(t1 string, t2 string) (str1 string, str2 string){
				t1 = strings.ToLower(t1)
				t2 = strings.ToUpper(t2)
				return t1, t2
}