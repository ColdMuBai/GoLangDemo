// package main
package main

import (
	"fmt"
	"os"
)

import "rsc.io/quote"



func main(){
	fmt.Println(quote.Go())

	var str,s string
	str = "Hello GoLang"
	s = ""

	// for i := 1; i < len(os.Args); i++{
	// 	str += s + os.Args[i]
	// 	s = ""
	// 	fmt.Println("Args[",i,"]=",os.Args[i]) 
	// }

	// fmt.Println(str)
	i := 0
	for _,arg := range os.Args[0:] {
		str += s + arg
		s = " "
		fmt.Println("idnex = ",i,"Value = ",arg)
		i++
	}

	fmt.Println(str)

}