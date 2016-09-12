package main

import (
	"fmt"
)

// 
var varType1 string
var varType2 int
//
var varType3, varType4 = "uno", "dos"
//
var (
	varType5 = "cinco" 
	varType6 = "seis"
)
//
const (
	varType7 = "Test7"
	varType8 = "Test8"
)


func main() {
	varType9 := "Test9"
	fmt.Println(varType1)
	fmt.Println(varType2)
	fmt.Println(varType3, varType4)
	fmt.Println(varType5, varType6)
	fmt.Println(varType7, varType8)
	fmt.Println(varType9)
}
