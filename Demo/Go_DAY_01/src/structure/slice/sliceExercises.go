package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	_ "strings"
)

func main() {
	// b := FindDigits("./slice.txt")
	// fmt.Printf("find result : %s",b)

	// s1,s2 := test01("生化危机8:村庄 豪华版",20)
	// fmt.Printf("Game: %v\t type: %v",s1,s2)

	// test02()
	// test03()
	// test04()
	// test05()
	test06()

}

var digitRegexp = regexp.MustCompile("[0-9]{3,}")

func FindDigits(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}

func test01(s string, i int) (s1 string, s2 string) {
	s1, s2 = s[:i], s[i:]
	return s1, s2
}

func test02() {
	str := "Snowflakes are kisses from heaven."
	var str01 string = str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str01)
}

func test03() {
	// sl := []byte("Snowflakes are kisses from heaven.")
	// var s []byte
	// for i := 0;i < len(sl);i++ {
	// 	s = append(s,sl[len(sl)-1-i])
	// }
	// fmt.Println(string(s))

	var sl = []rune("这是一段正序的string 字符串!")
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	fmt.Println(string(sl))
}

func test04() {
	var s = []byte(`Access InterNet mmmubb`)
	var s1 = make([]byte, 0)
	//将于前一个字符相同的添加到数组中
	// var c byte
	// for _,v := range s {
	// 	if(c == v){
	// 		if len(s1) > 0 && v == s1[len(s1)-1] {
	// 			s1 = append(s1,v)
	// 		} else {
	// 			s1 = append(s1,c,v)
	// 		}
	// 	} else {
	// 		c = v
	// 	}
	// }

	//将与前一个字符不相同的添加到数组中
	var c byte
	for _, v := range s {
		if v != c {
			if len(s1) > 0 && s1[len(s1)-1] != c {
				s1 = append(s1, c, v)
			} else {
				s1 = append(s1, v)
			}
			c = v
		}
	}

	fmt.Printf("same character : %s", s1)
}

func test05() {
	var s = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	var temp int
	var index int
	for i := range s {
		temp = s[i]
		for j, v := range s[i:] {
			if v <= temp {
				temp = v
				index = j + i
			}
		}
		if index != i {
			s[index], s[i] = s[i], temp
		}
		// fmt.Printf("new slice : %v\n",s)

	}
	fmt.Printf("s : %v", s)
}

func test06() {
	var intArr = []int{2, 4, 6, 8, 10}
	intArr = MapFunc(intArr)
	fmt.Printf("%v",intArr)
}

func MapFunc(s []int) []int {

	for i, v := range s {
		s[i] = v * 10
	}
	return s
}
