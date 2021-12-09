package main

import (
	"fmt"
)

func f() {}


func main() {
	if x := f();x ==0 {
		fmt.Println(x)
	} else if y := g(x);x ==y {
		fmt.Println(x,y)
	} else {
		fmt.PrintLn(x,y)
	}
	fmt.Println(x,y)

}