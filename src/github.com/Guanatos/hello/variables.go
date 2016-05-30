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


func main() {
	fmt.Println(varType1)
	fmt.Println(varType2)
	fmt.Println(varType3, varType4)
	fmt.Println(varType5, varType6)
}
