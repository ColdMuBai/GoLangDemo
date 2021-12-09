package main

import(
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main(){
	for _,arg := range os.Args[1:] {
		t,err := strconv.ParseFloat(arg,64)
		if err != nil {
			fmt.Fprintf(os.Stderr,"CF : %v\n",err)
			os.Exit(1)
		}
		printTemp(t)
		printUnit(t)
	}
	
}

func printTemp(t float64){
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s,\t%s = %s\n",
		f,tempconv.FToC(f),c,tempconv.CToF(c))
}

func printUnit(t float64){
	m := tempconv.M(t)
	f := tempconv.FP(t)
	k := tempconv.KG(t)
	b := tempconv.LB(t)
	fmt.Printf("输入值为: %v\n%v =%v \n%v = %v \n%v = %v \n%v  = %v",
		t,m,tempconv.MToFP(m),f,tempconv.FPToM(f),k,tempconv.KGToLB(k),
		b,tempconv.LBToKG(b))
}