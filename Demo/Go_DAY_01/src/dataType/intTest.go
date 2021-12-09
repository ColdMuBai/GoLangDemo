package main

import (
	"fmt"
	"strconv"
	"math"
)


func main(){
	// FloatTest()
	// intIota()
	rangeCirculation()
	
}




type student struct {
    id   string
    name string
    sex  rune
}

func rangeCirculation(){
    ss := []student{
        {"1", "张三", '男'},
        {"2", "李四", '女'},
        {"3", "王五", '女'},
        {"4", "赵六", '男'},
        {"5", "孙七", '女'},
    }
    for i,s := range ss {
        fmt.Printf("i= %v \t%p\t s = %v\n",i,&s,s.name)
    }
}

func FloatFmt(){
	for x := 0;x < 8; x++{
		fmt.Printf("x = %d e^x = %8.3f\n",x,math.Exp(float64(x)))
	}
}

func FloatTest(){
	//格式化很小或者很大的数时 使用 %g 格式化输出
	const Avogadro = 6.02214129e23    //阿伏伽德罗常数
	const Planck = 6.62606957e-34     //普朗克常数
	fmt.Println(Planck)		//效果等同于%g的打印效果
	fmt.Printf("Avogadro = %f  \nPlanck = %f\n",Avogadro,Planck)
	fmt.Printf("Avogadro = %g  \nPlanck = %g",Avogadro,Planck)
}


func IntFlow(){
	var u uint8 = 255
	fmt.Println(u)
	fmt.Println(u+1)
	fmt.Println(u*u)
	i,err  := strconv.ParseInt("256",10,8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T : %v",i)
}

func intIota(){
	// const (
	// 	a = iota
	// 	b 
	// 	c
	// 	d = 500
	// 	e
	// )
	const (
		a = 1
		b 
		c
		d = 500
		e
	)
	fmt.Printf("a= %d \nb = %d \nc= %d \nd= %d \ne= %d",a,b,c,d,e)

	const (
		Apple,Banana = iota +1,iota +3 
		Cherimoya,Durian
		Elderberry,Fig
	)

	fmt.Printf("\nApple = %v \nBanana = %v \nCherimoya = %v \nDurian = %v \nElderberry = %v \nFig = %v",Apple,Banana,Cherimoya,Durian,Elderberry,Fig)

	const (
		Open = 1 << iota 
		Close
		Pending
	)

	fmt.Printf("\nOpen = %v \nClose = %v \nPending = %v",Open,Close,Pending)

	// const (
	// 	_ = iota 	//使用_忽略不需要的 iota
	// 	KB = 1 << (10 * iota)          // 1 << (10*1)
	// 	MB                             // 1 << (10*2)
	// 	GB                             // 1 << (10*3)
	// 	TB                             // 1 << (10*4)
	// 	PB                             // 1 << (10*5)
	// 	EB                             // 1 << (10*6)
	// 	ZB                             // 1 << (10*7)
	// 	YB                             // 1 << (10*8)
	// )
	// fmt.Printf("KB = %v \nMB = %v \nGB = %v \nTB ",
	// 	"= %v \nPB = %v \nEB = %v \nZB = %n \nYB = %v",
	// 	KB,MB,GB,TB,PB,EB,ZB,YB)
}

