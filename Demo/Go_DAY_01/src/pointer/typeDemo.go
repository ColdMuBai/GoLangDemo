package main

import (
	"fmt"
)


var (
	a = 10.2
	b = 35
	c = "这是一个批量设置变量中的string类型变量"
)

type Celsius float64 	//摄氏温度
type Fahrenheit float64 //华氏温度

// const(
// 	AbsoluteZeroC Celsius = -273.15 //绝对零度
// 	FreezingC Celsius = 0 	//冰点温度
// 	BoillingC Celsius = 100 //沸水温度
// )


func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5+32)}
func FToC(f Fahrenheit) Celsius {return Celsius((f-32)*5 /9)}

func main(){
	// fmt.Printf("a= %f\nb= %d\nc= %s\n",a,b,c)
	// fmt.Printf("%g\n",BoillingC-FreezingC)
	// boilingF := CToF(BoillingC)
	// fmt.Printf("%g\n",boilingF-CToF(FreezingC))
	//fmt.Printf("%g\n",boilingF-FreezingC)

	//-------------------------------------------------------------------
	// var c Celsius
	// var f Fahrenheit
	// fmt.Println(c ==0.00)		//"true"
	// fmt.Printf("C type of %T\n",c)
	// fmt.Println(f >=0)		//"true"
	// fmt.Printf("c= %v;f= %v",c,f)
	// //fmt.Println(c ==f)		//"compile error : type mismatch"
	// fmt.Println(c == Celsius(f)) //"true"


	ToString();
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃",c)
}

func ToString(){
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n",c)
	fmt.Printf("%s\n",c)
	fmt.Println(c)
	fmt.Printf("%g\n",c)
	fmt.Println(float64(c))
}
