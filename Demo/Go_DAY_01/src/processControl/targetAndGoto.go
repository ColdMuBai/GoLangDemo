package main

import (
	"fmt"
)

func main() {
	// target01()
	// gotoTest()
	gotoErr()
}

func target01(){
	TARGET: 
	for i := 0; i <=5 ; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				fmt.Println()
				continue TARGET
			}
			fmt.Printf("i is: %d, and j is: %d\n",i,j)
		}
	}
}


func gotoTest(){
	i := 0 
	GOTOTS:
		print(i)
		i++
		if i == 5{
			return
		}
		goto GOTOTS
}

func gotoErr(){
	i := 0
	for { //since there are no checks, this is an infinite loop
	    if i >= 3 { break }
	    //break out of this for loop when this condition is met
	    fmt.Println("Value of i is:", i)
	    i++
	}
	fmt.Println("A statement just after for loop.")
}