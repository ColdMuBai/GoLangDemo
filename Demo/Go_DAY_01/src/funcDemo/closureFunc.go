package main
//匿名闭包函数

import (
	"fmt"
	"strings"
)

func main () {
	// f1 := Add01()
	// fmt.Printf("function Add2 for gives%v\n",f1(9))

	// var f = AddrPlus01()
	// fmt.Print(f(1),"-")
	// fmt.Print(f(20),"-")
	// fmt.Print(f(300),"-")

	// f := fibFunc()
	// arr := make([]int,0)
	// for i := 0; i < 100; i++ {
	// 	arr = append(arr,f())
	// }
	// fmt.Printf("%v\n",arr)
	// fmt.Printf("切片长度为%d,容量为%d\n",len(arr),cap(arr))

	addGo := AutoAddSuffix(".go")
	addHtml := AutoAddSuffix(".html")

	fmt.Printf("%s",addGo("这是一个GoLang文件!"))
	fmt.Printf("%s",addHtml("这是一个HTML文件!"))

}

func Adder(a int) func(b int) int {
	return func (b int) int {
		return a+b
	}
}

func Add01() func (b int) int {
	return func(b int) int{
		return b + 2
	}
}
func AddrPlus01() func(int) int {
	var x int
	return func( delta int) int {
		x += delta
		return	x
	}
}

func FibFunc() func () int{
	 num1,num2 := 0,1
	 return func () int {
	 	num1,num2 = num2,num1+num2
	 	return num2
	 }
}

func AutoAddSuffix(suffix string) func(string) string{
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}	

}