package main

import (
	"fmt"
	"math/big"
)

func main(){
	// var arrs  = []int{3,5,6,7,1,5,9,11,0,2,45}
	// var arrs1 = []int{65,642,5,2,1,0,65,-50,87,81,0,452,3}
	// quick(&arrs,0,len(arrs)-1)
	// quick(&arrs1,0,len(arrs)-1)
	// fmt.Printf("arrs: %v\narrs1: %v",arrs,arrs1)

	// result := Factoral(30)
	// fmt.Printf("result : %v \tType: %[1]T \n %#v",result,*result)
	// fmt.Printf("%T,%[1]v",big.NewInt(1))

	// PrintRangeNum(10)

	PrintFFib()

}


//获取基准值
func QuickSort(arrs *[]int,start int,middel int) int{
	arr := *arrs
	temp := arr[start]
	for start < middel{
		for  arr[middel] >= temp && start < middel{
			middel--
		}
		arr[start] = arr[middel]
		for arr[start] <= temp && start < middel{
			start++
		}
		arr[middel] = arr[start]
	}
	arr[start] = temp
	*arrs = arr
	return start
}
//递归调用直到不能再次拆分
func quick(arr *[]int,start int,middel int){
	if(start < middel){
		// fmt.Printf("start: %v,middel: %v",start,middel)
		index :=  QuickSort(arr,start,middel)
		// fmt.Printf("index : %v,arr : %v\n",index,*arr)
		quick(arr,start,index-1)
		quick(arr,index+1,middel)
	}
}

//计算前n个整数的阶乘
func Factoral(n int) (product *big.Int){
	if(n ==1){
		return big.NewInt(1)
	}
	current := big.NewInt(int64(n))
	product = current.Mul(Factoral(n-1),current)
	// i := big.NewInt(1)
	// product = i.Mul(big.NewInt(10),i)
	return 
}

//递归输出10 ~ 1
func PrintRangeNum(num int){
	if(num >=1){
		println("当前Number值为 : ",num)
		PrintRangeNum(num -1)
	}
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
	for i := 0; i <=10 ;i++ {
		fmt.Printf("Fibonacci(%v) is : %[1]v\n",i,Fibonacci(i))
	}
}