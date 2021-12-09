package main

import (
	"fmt"
	"time"
)

//声明斐波那契数列阈值
const LIM = 40
//声明存储斐波那契数列的数组
var fib [LIM]uint64


func main (){
	// caclTime()
	CacheFib()
	// PrintFFib()
}

func caclTime() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		fmt.Printf("循环%v次后结束\n",100-i)
	}
	delta := time.Now().Sub(start)
	fmt.Printf("forLoop took is amount of time %s\n",delta)

}

func fibonacci(n int) (res uint64) {
	if fib[n] != 0 {
		res = fib[n]
		return
	}
	if n <=1 {
		res = 1
	} else {
		res = fibonacci(n-1)+fibonacci(n-2)
	}
	fib[n] = res
	return
}



func CacheFib(){
	var result uint64 = 0
	start := time.Now()

	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n",i,result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("The consumption of time is %s",delta)
}



//计算斐波那契数列
func Fibonacci(n int) (res int){
	if (n <=1){
		res =1
	} else {
		res = Fibonacci(n -1) +Fibonacci(n -2)
	}
	return
}

func PrintFFib(){
	start := time.Now()
	for i := 0; i <=LIM ;i++ {
		fmt.Printf("Fibonacci(%v) is : %[1]v\n",i,Fibonacci(i))
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("The consumption of time is %s",delta)
}