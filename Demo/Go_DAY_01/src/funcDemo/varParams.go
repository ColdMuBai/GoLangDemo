package main

import (
	"fmt"
	"regexp"
)

func main(){
	x := findMinNumber(6,3,2,5,9,2,50)
	fmt.Printf("the minNumber is : %d\n",x)
	
	slice := []int{9,3,6,8,5,1,2,-10001}
	//slice... 拆解切片中的元素
	x = findMinNumber(slice...) 
	fmt.Printf("the minNumber is : %d\n",x)

	// varargsPrint()
}

func findMinNumber(s ...int) (min int){
	if len(s) == 0 {
		return 0
	}

	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return
}


func varargsPrint(s ...string){
	const str string = "这是一段预留的待处理文本"

	for _, item := range regexp.MustCompile(`待处理`).ReplaceAllString(str,"处理后") {
		fmt.Printf("%c\n",item)
	 } 

}