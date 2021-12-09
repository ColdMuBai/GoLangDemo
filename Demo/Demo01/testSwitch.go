package main

import(
	"fmt"
)
func main(){
	var i int= Signum(10)
	fmt.Println(i)
}

func Signum(x int) int {
	switch{
		case x > 0:
			return +1
		default:
			return 0
		case x <0 :
			return -1
	}
	
}