package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


func main(){
	// BasicSwitch()
	// FallSwitch()
	// CaseMoreVal()
	// InitSwitch()
	// VerdictSeason()
	// ScanContent()
	BufioRead()

}


func BasicSwitch(){
	var i int = 100
	// var i8 int8 = 100
	// var i32 int32 = 100

	//类型不一致时无法进行比较,编译无法
	// switch i {
	// 	case i :
	// 		fmt.Printf("target type: %T val: %[1]v",i)
	// 	case i8 :
	// 		fmt.Printf("target type: %T val: %[1]v",i8)
	// 	case i32 :
	// 		fmt.Printf("target type: %T val: %[1]v",i32)
	// }
	switch {
		case i >100 :
			fmt.Printf("target type: %T val: %[1]v",i)
		case i/3 >35:
			fmt.Printf("target type: %T val: %[1]v",i/3)
		case i%2 ==0 :
			fmt.Printf("target type: %T val: %[1]v",i%2)
	}
}

func FallSwitch(){
	var flag int = 10000
	switch flag {
		case 0 : 
			fmt.Printf("flag的值为 : %v\n",flag)
			flag++;fallthrough
		case 10 : 
			fmt.Printf("flag的值为 : %v\n",flag)
			flag += 100
			fallthrough
		case 99 :
			fmt.Printf("flag的值为 : %v\n",flag)
			flag -= 1
			fallthrough
		case 100 :
			fmt.Printf("flag最终的值为 : %v\n",flag)
		default :
			fmt.Println("flag没有匹配到任何值")
	} 
}

func CaseMoreVal(){

	var num = 66

	switch (num) {
		case 70,69,68 :
			fmt.Println("num equal 70 or 69 or 68")
		case 67,66,65 :
			fmt.Println("num equal 67 or 66 or 65")
		default :
			fmt.Println("It's not equal none")
	}
}


func InitSwitch(){
	x := [15]int{1,2,3,4,5,6}
	y := [16]int{1,2,3,4,5,6}
	var t int
	switch a, b := x[5],y[4]; {
	    case a < b: t = -1
	    // case a == b,true: t = 0
	    //同时case多个条件时使用,进行分割
	    case a == b: t = 0
	    case a > b: t = 1
	}
	fmt.Printf("final t : %v",t)
}

func VerdictSeason(season int8){
   switch season {
   		case 1,2,3:
   			fmt.Printf("%v月份是春季\n",season)
		case 4,5,6:
			fmt.Printf("%v月份是夏季\n",season)
		case 7,8,9:
			fmt.Printf("%v月份是秋季\n",season)
		case 10,11,12:
			fmt.Printf("%v月份是冬季\n",season)
		default:
			fmt.Println("月份都输不明白还想知道季节?")
   }
}


func ScanContent(){
	for{
		fmt.Scan()
		//从stdin中读取内容直到遇到换行符时停止
		input,err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		readStr := strings.TrimSpace(input)
		if readStr == "q" {
				os.Exit(0)
				return
		}
		inputRead,err := strconv.Atoi(readStr) 
		if err != nil {
			fmt.Println("输入的内容非月份数")
		}
		VerdictSeason(int8(inputRead))
	}
	
}


func BufioRead(){
	fmt.Println("内容读取开始,'退出程序'请输入Exit")
	fmt.Println("请输入需要判断季节的月份")
	input := bufio.NewScanner(os.Stdin)

	for {
		input.Scan()
		inputRead := strings.TrimSpace(input.Text())
		switch inputRead {
			case "": continue
			case "Exit": os.Exit(0);return
			default : 
				month,err := strconv.Atoi(inputRead) 
				if err != nil {
					fmt.Println("输入的内容非月份数")
					break
				}
				VerdictSeason(int8(month))
		}
	}

}