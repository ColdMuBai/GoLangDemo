package main

import (
	"fmt"
	"os"
	"runtime"
	"math"
)

var a = "G"

func main (){
	// PrintOsInfo()
	// runtimeTest()
	casInt,err :=  Uint8FromInt(122)
	if err != nil {
		fmt.Printf("switch error , message = %v\n",err)
	}
	fmt.Printf("casInt type is %T  val= %v\n",casInt,casInt)

	casFloat := IntFromFloat64(65535.1024)
	fmt.Printf("casInt type is %T  val= %v\n",casFloat,casFloat)

}

func Uint8FromInt(n int) (uint8 ,error){
	if 0<= n && n <= math.MaxUint8 {
		return uint8(n),nil
	}
	return 0,fmt.Errorf("%d 超出了 uint8 的存储范围",n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole,fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("g 超出了int32类型的存储范围",x))
}

func runtimeTest(){
	var goos string = runtime.GOOS
	fmt.Printf("The operating syxtem is: %s\n",goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n",path)
}

func PrintOsInfo(){
	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
		GPATH = os.Getenv("GOPATH")
	)	
	fmt.Printf("HOME = %v \nUSER = %v \nGPATH = %v",HOME,USER,GPATH)
}