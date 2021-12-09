package main	

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	RandomTest01()
	fmt.Println()
	RandomTest02()
	RandomTest03()
}



func RandomTest01(){
	for i := 0; i < 10; i++ {
		a := rand.Int()	
		if((i+1)%3 == 0){
			fmt.Printf("%d\n",a)
			continue
		}
		fmt.Printf("%d\t",a)
	}
}

func RandomTest02(){
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		if((i+1)%3 == 0){
			fmt.Printf("%d\n",r)
			continue
		}
		fmt.Printf("%d\t",r)
	}
}

func RandomTest03(){
	fmt.Println()
	times := int64(time.Now().Nanosecond())
	rand.Seed(times)
	for i := 0; i < 10; i++ {
	 	if((i+1)%3 == 0){
			fmt.Printf("%2.2f\n",100*rand.Float32())
			continue
		}
		fmt.Printf("%2.2f\t",100*rand.Float32())
	 } 
}
