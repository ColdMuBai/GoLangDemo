package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)


func main() {

	RepeatDemo()
}


func RepeatDemo() {
	var str string= "将s的每一个unicode码值r都替换为mapping(r),"
	fmt.Println(str[:])
	s := strings.Map(FormatUnitCode,str)
	fmt.Println(s)
}

func FormatUnitCode(r rune) rune{

	if r > utf8.RuneSelf {
		return '?'
	}
	return r

}