package main

import (
	"time"
	"fmt"
)

func main(){
	// timeTest()
	// timestapTest()
	compareTime()
	// timeTickTest()
	// timeAfterFuncTest()
	// timeTickTest()
}

//时间包time基本使用
func timeTest(){
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%02d-%02d-%02d\n",now.Year(),now.Month(),now.Day())
	fmt.Printf("%02d-%02d-%02d\n",now.Hour(),now.Minute(),now.Second())
	//时间格式化成字符串只认GO语言的诞生日期
	fmt.Printf("TimeFormat : %v\t FormatType : %[1]T\n",now.Format("2006-01-02 15:04:05"))	
	fmt.Printf("TimeFormat : %v\t FormatType : %[1]T\n",now.Format("2006年1月2日 15小时4分5秒"))
	//日期字符串格式化为时间
	t1,err := time.Parse("2006-01-02 15:04:05","2021-11-24 09:21:50")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("TimeFormat : %v\t FormatType : %[1]T\n",t1)
	t2,err := time.Parse("2006年1月2日 15小时4分5秒","2021年11月24日 09小时21分50秒")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("TimeFormat : %v\t FormatType : %[1]T\n",t2)
}

//unix时间格式化
func timestapTest(){
	/**
	 * time.Now().Unix() 获取当前时间的时间戳
	 *  time.Unix(timestap,0)将时间戳转化为时间
	 *  Unix 从January 1, 1970 UTC至该时间的秒数和纳秒数）。
	 *  nsec 的值在 [0, 999999999] 范围外是合法的。
	 * func Unix(sec int64, nsec int64) Time
	 * sec 时间戳 ,秒级 
	 * nesc 纳秒级时间(一般设置为0)
	 * 两者皆可进行时间偏移使用
	 */
	tm := time.Unix(time.Now().Unix(),1e9*60*60)
	fmt.Printf("timestap : %v\t type : %[1]T\n",tm)
	tmFormat := tm.Format("2006-01-02 15:04:05")
	// tmFormat := tm.Format("2006年01月02日 15时04分05秒")
	fmt.Printf("tmFormat : %v\t type : %[1]T\n",tmFormat)

	/**
	 * 组装时间
	 * <年 月 日 时 分 秒 纳秒 时区>
	 * time.Local 返回本地时区
	 */
	var timeNow time.Time= time.Now()
	timeEnd := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 0, time.Local)
	fmt.Printf("timeEnd : %v\t %[1]T\n",timeEnd)
}

/**
 * 时间类型比较
 * 逻辑运算符中只有 == 被定义过
 * time 包下的比较方法
 * Equals	是否相等
 * Before	t1 是否在 t2 之前
 * After	t1 是否在 t2 之后
 * 
 * 时间类型加减
 * 
 */ 
func compareTime(){
	var t1 time.Time= time.Now()
	t2,_ := time.Parse("2006-01-02 15:04:05","2022-11-24 10:15:36")
	fmt.Printf("t1 == t2 ? %[1]v",t1 == t2)
	fmt.Printf("t1 Equal t2 ? %[1]v \nt1.Before(t2) ? %[2]v \nt1.After(t2) ? %[3]v\n",t1.Equal(t1),t1.Before(t2),t1.After(t2))
	/**
	 * 
	 * func (t Time) Date() (year int, month Month, day int)
	 * func (t Time) Clock() (hour, min, sec int)
	 */ 
	y,M,d := t1.Date()
	fmt.Printf("Current Date : %v-%d-%v\n",y,M,d)
	h,m,s := t1.Clock()
	fmt.Printf("Current Clock :  %v:%v:%v\n",h,m,s)
	/**
	 * 返货时间店 t + d
	 * 可选择的计数单位 
	 * Duration类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年。
	 * const (
	 *	  Nanosecond  Duration = 1
	 *	  Microsecond          = 1000 * Nanosecond
	 *	  Millisecond          = 1000 * Microsecond
	 *	  Second               = 1000 * Millisecond
	 *	  Minute               = 60 * Second
	 *	  Hour                 = 60 * Minute
	 *	)
	 * 
	 */ 
	fmt.Printf("Add Hour: %[1]v\nSub Hour : %[2]v\n",t1.Add(60 * time.Minute),t2.Sub(t1))
	fmt.Printf("AddDate 1年1月3天: %v\n",t1.AddDate(1,1,3))
	fmt.Printf("Round time t1 : %v\n",t1.Round(1 * time.Hour))
	fmt.Printf("Truncate time t1 : %v\n",t1.Truncate(1 * time.Hour))

}

//time.After 等待时间d后发送当前时间
func timeAfterTest(){
	fmt.Println(time.Now())
	//返回一个channel 10s后向channel 发送创建时的时间
	c := time.After(10 * time.Second)
	t := <-c
	fmt.Println(t)
	tm := time.NewTimer(10 * time.Second) //NewTimer 返回 Timer类型
	t = <-tm.C  //Timer结构中有一个channel,C,10秒后将当前时间发送到C中
	fmt.Println(t)
}

//time.AfterFunc 等待时间d然后调用函数f
func timeAfterFuncTest(){
	fmt.Println(time.Now())
	time.AfterFunc(10 * time.Second,Test)
	var str string
	fmt.Scan(&str) //等待用户输入,防止线程提前结束
}	

func Test(){
	fmt.Printf("welcome come GoLang : 这是 %v 设置的延时10s的信息!",time.Now())
}


//time.Tick 重复执行的定时任务
func timeTickTest(){

	// c := time.Tick(2 * time.Second)
	// var i int = 1
	// for current := range c {
	// 	i++
	// 	fmt.Println(current)
	// 	if(i >=10){
	// 		return
	// 	}
	// }
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		Test()
	}

}		

func timeTickerTest(){
	//TODO
}