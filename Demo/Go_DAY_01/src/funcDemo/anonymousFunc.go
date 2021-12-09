package main

import (
	"fmt"
)


func main() {
	AnoFunc01()

	fv := func(){print("Hello World!")}
	fv()
	fmt.Printf("fv type : %T fv value : %[1]v",fv)
}


func AnoFunc01() {
	for i := 0; i < 10; i++ {
		// f := func(i int) { fmt.Printf("循环%d次",i)}
		// f(i)
		func(i int) { fmt.Printf("循环%d次\n",i)}(i)
		// fmt.Printf("func f is of type %T and has value %v\n",f,f)
	}

}