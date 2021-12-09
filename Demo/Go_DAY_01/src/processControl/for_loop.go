package main

import (
	"fmt"
)


func main() {
	// test01()
	// test02()
	// test03()
	// test04()
	// test05()
	// rangeLoop01()
	t1()
	
}
func t1(){
	for i, j, s :=0,5,"a"; i <3&& j <100&& s !="aaaaa"; i, j,
    s = i+1, j+1, s +"a"{
    fmt.Println("Value of i, j, s:", i, j, s)
	}
}



func test01(){
	for i := 0; i < 15; i++ {
		fmt.Printf("当前计数器i的值为 %d\n",i)
	}
		
}

func test02(){
	// for i := 1; i <= 25; i++ {
	// 	for k := 0; k < i; k++ {
	// 		fmt.Print("G")
	// 	}
	// 	fmt.Println()
	// }
	var temp string = "G"
	for i := 1; i <= 25; i++ {
		fmt.Println(temp)
		temp += "G"
	}
}

func test03(){
	for i := 0; i <= 10; i++ {
		fmt.Printf("current i= %v ^i = %12b\n",i,^i)
		// var f float32= float32( i/2.0)
		// fmt.Printf("%2.2fxxxx%[1]f\n",float32( i/2.0))
		// fmt.Printf("%3.6dxxxx%[1]d\n",i+12000)
	}
}

func test04(){

	for i := 1; i <= 100; i++ {
		var formatStr = "%v\t"
		var finalVal string
		switch {
			case i%10 == 0: 
				formatStr += "\n";
				finalVal = "Buzz";
				if(i % 15 ==0){
					finalVal = "FizzBuzz"
				}
			case i%15 == 0: finalVal = "FizzBuzz"
			case i%3 == 0: finalVal = "Fizz"
			case i%5 == 0: finalVal = "Buzz"
			default:
				finalVal = fmt.Sprintf("%d",i)
		}
		fmt.Printf(formatStr,finalVal)
	}
}

func test05(){
	for i := 0; i < 10; i++ {
		for k := 0; k <20; k++ {
			if(i == 0 || i == 9 || k == 0 || k == 19){
				fmt.Print("* ")
				continue
			}
			fmt.Printf("  ")
		}
		fmt.Println()
	}
}
func rangeLoop01(){
	str := "GO     is a beautiful language"
	const subStr = "Lang" 
	fmt.Printf("The length of str is: %d\n",len(str))
	for pos, char := range str {
		if(pos>1 && pos<6){
			char = rune(subStr[pos-2])
		}
		fmt.Printf("Character on position %d is: %c \n",pos,char)
	}
	str2 := "Chinese 路过一只猫"
	fmt.Printf("\nthe length of str2 is: %d\n",len(str2))
	for pos,char := range str2 {
		fmt.Printf("Character %c starts position %d\n",char,pos)
	}
	fmt.Println("\nindex int(rune) rune  char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d	%[2]d	%[2]U	'%[2]c' %[3]x\n",index,rune,[]byte(string(rune)))
	}

}