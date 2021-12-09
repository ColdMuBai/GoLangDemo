package tempconv

import "fmt"

//定义长度单位
type FP float64
type M float64

//定义重量单位
type KG float64
type LB float64


//默认打印 String()方法
func (f FP) String() string{return fmt.Sprintf("%g fp",f)}

func (m M) String() string{return fmt.Sprintf("%g m",m)}

func (k KG) String() string{return fmt.Sprintf("%g kg",k)}

func (b LB) String() string{return fmt.Sprintf("%g lb",b)}

//单位换算方法
func MToFP (m M) FP {
	return FP(m/0.3048)
}

func FPToM (f FP) M {
	return M(f*0.3048)
}

func KGToLB (k KG) LB {
	return LB(k/0.45359237)
}

func LBToKG (b LB) KG {
	return KG(b *0.45359237)
}

