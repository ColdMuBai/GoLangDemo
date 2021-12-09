package main

import (
	"fmt"
)


func main() {

}

func test01() {
	var mapList map [string]int
	var mapAssigned map[string]int

	mapList = map[string]int{"one":1,"two":2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapList

	mapCreated["key1"] = 26.8
	mapCreated["key2"] = 28.8

	fmt.Print("mapList : %v",mapList)
	fmt.Print("mapCreated : %v",mapCreated)
	fmt.Print("mapAssigned : %v",mapAssigned)
}