package main

import (
	"fmt"
)


func main() {
	arr1 := [5]float64{1.1,3.3,5.5,7.7}
	sum := transferArr(&arr1)
	fmt.Println("arr1 lengths is ",len(arr1))
	fmt.Printf("Current Arrays is %[2]v\nCurrent Arrays is %[1]v",arr1,sum)
}

func arrInit(){
	a := [...]string{"a","b","c","d"}
	var s []string
	s = a[:]
	fmt.Printf("[...]string{ } type is %T\tval : %[1]v\n",a)
	fmt.Printf("slice type is %T\tval: %[1]v",s)
}

func transferArr(arr *[5]float64) (sum float64){
	for i, v := range arr {
		fmt.Printf("Current num is %f\n",v)
		sum += v
		arr[i] = sum
	}
	return
}