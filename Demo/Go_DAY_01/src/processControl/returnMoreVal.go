package main

import (
	"fmt"
	"strconv"
	"os"
)


func main() {
	readMkDocs()
}


func readMkDocs(){
	f,err := os.Open("./GoLang学习笔记.md")
	//defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	defer f.Close() 
	if err !=nil {
		//主动触发宕机
		panic(err)
	} else {
		buffer := make([]byte,20)
		//func (f *File) Read(b []byte) (m int,err errpr)
		//使用read方法,将文件内容读入到buffer切片中
		for {
			// length,err := f.Read(buffer)
			_,err := f.Read(buffer)
			if err != nil {
				panic(err)
			} else {
				// fmt.Println("读取了",length,"字节的内容")
				//fmt.Println(string(buffer)) //输出时会乱码
				fmt.Printf("%v",string(buffer)) // 以读取内容变量的自然形式输出
			}
		}
	}

}


func test01(){
	var orig string = "MuBai"
	var newS string

	//IntSize是int或uint类型的字位数
	fmt.Printf("The size of ints is : %d\n",strconv.IntSize)

	an,err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("error : %s\n",err)
		//在错误发生的同时终止程序的运行,下方代码均不执行 1 -->退出状态码;可被脚本获取
		os.Exit(1)
		fmt.Printf("error : %s\n",err)
		// return 
	}
	fmt.Printf("string to Integer %d\n",an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("Integer to string :%q",newS)
}