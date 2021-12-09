package main

import (
	"fmt"
)


func main(){
	PtrTest01()
}

func PtrTest01(){
	var i =5
	fmt.Printf("i= :%d ,i->pointer: %p\n",i,&i)
	var intP *int = &i
	fmt.Printf("intp : %v\t val : %v pointer: %p\n",intP,*intP,&intP)

	var s = "init Strings"
	var p *string = &s
	*p = "alert String"
	fmt.Printf("s address: %p\tp address: %p\n",p,&p)
	fmt.Printf("*p val: %2v type: %[1]T\n",*p)
	fmt.Printf("s val: %2v type: %[1]T",s)
}