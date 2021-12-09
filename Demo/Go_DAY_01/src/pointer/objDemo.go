package main

import (
	"fmt"
)


func main(){
	// p := new(string)
	// *p = "string变量"
	// fmt.Println(*p)

	// var x,y = 9,6
	// x,y = y,x
	// fmt.Printf("x = %v,y= %v\n",x,y)
	// var arr = [...]int{1,3,5,7,9}
	// arr[0],arr[3] = arr[1],arr[4]
	// fmt.Println(arr)
	// fmt.Println(arr[0],arr[3])

	fmt.Println(gcd(6,9))
	fmt.Println(fib(9))

}


func gcd(x,y int) int{
	for y !=0 {
		x,y = y, x%y 			
	}
	return x
}


func fib(n int) int {
	x,y := 0,1
	for i := 0; i< n-1;i++{
		x,y = y,x+y   		
	}
	return x
}