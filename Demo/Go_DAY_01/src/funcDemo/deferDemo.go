package main

import (
	"fmt"

)

func main (){
	printI()

}


func printI(){
	i := 0
	defer fmt.Println("i=",i)
	i = 10
	fmt.Println("this function is over")
}