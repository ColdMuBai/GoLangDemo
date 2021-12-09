package main

import (
	"fmt"
	"runtime"
)
var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."

func main(){
	fmt.Println(prompt)
}

func osTypeTest(){
	const osType = runtime.GOOS
	fmt.Println(osType)
	if osType == "windows" {
		fmt.Println("运行系统为Windows")
	} else {
		fmt.Println("运行系统不是windows,而是",osType)
	}
}
//演示如何根据操作系统来决定输入结束的提示：
func init(){
	if runtime.GOOS == "windows" {
         prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")        
     } else { //Unix-like
         prompt = fmt.Sprintf(prompt, "Ctrl+D")
     }
}
