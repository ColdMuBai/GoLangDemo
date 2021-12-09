package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ","separator")


func main(){
	// var x,y int 
	// fmt.Printf("x =%T\ny =%T\n",&x,&y)
	// fmt.Println(&x == &x, &x == &y, &x == nil)
	// fmt.Println(f() == f())

	// v := 1
	// fmt.Printf("v = %v\t&v = %v\n",v,&v)
	// incr(&v)
	// fmt.Printf("v = %v\t&v = %v\n",v,&v)
	// fmt.Println(incr(&v))

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}

}

func f() *int{
	v := 1
	fmt.Println("&v=",&v)
	return &v
}

func incr(p *int) int{
	*p++  //此表达式只增加P指向的变量的值,并不改变P指针 !!!
	return *p
}
