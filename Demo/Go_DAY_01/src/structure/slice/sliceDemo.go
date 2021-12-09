package main

import (
	"fmt"
)

func main() {
	// ArrSlice()
	// slice01()
	// addByte()
	sliceLength()
}

func ArrSlice() {
	var arr [5]int = [5]int{1, 3, 5, 6, 7}
	fmt.Printf("arr val: %v\n", arr)
	s := arr[1:4]
	fmt.Printf("slice val: %v", s)
}

func slice01() {
	s := make([]byte, 5)
	s[0] = 'p'
	fmt.Printf("init slice Type : %T,len: %v\t cap : %v", s, len(s), cap(s))
}

func autoDilatation(s *[]byte, b byte) {
	arr := *s
	if float64(len(arr)) >= 0.75*float64(cap(arr)) {
		tempSlice := make([]byte, len(arr)+1, cap(arr)*2)
		copy(tempSlice, arr)
		// fmt.Printf("len is %v,%v\n",len(tempSlice),tempSlice)
		// fmt.Printf("len%v cap%v", len(tempSlice),cap(tempSlice))
		tempSlice[len(arr)] = b
		*s = tempSlice
	} else {
		*s = append(arr, b)
	}

}

func addByte() {
	byt := []byte{'a', 'c', 'd', 'e', 'f'}
	sl := make([]byte, 0, 2)
	for _, c := range byt {
		fmt.Printf("slice length: %v\t cap: %[2]v\t slice: %v\n", len(sl), cap(sl),sl)
		autoDilatation(&sl, c)
		fmt.Printf("slice length: %v\t cap: %[2]v\t slice: %v\n", len(sl), cap(sl),sl)
	}
}


func sliceLength(){
	var s = []int {1,3,5,7,9}
	fmt.Printf("%v",len(s[2:2]))
	fmt.Printf("%v",len(s[2:3]))
}