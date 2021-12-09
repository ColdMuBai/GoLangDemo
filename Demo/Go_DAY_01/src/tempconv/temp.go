package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const(
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC Celsius = 0 	//冰点温度
	BoillingC Celsius = 100 //沸水温度
)


func (c Celsius) String() string {
	return fmt.Sprintf("%g℃",c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F",f)
}

