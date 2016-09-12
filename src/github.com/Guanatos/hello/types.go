package main

import (
	"fmt"
	"strconv"
)

/*
bool
string
int
uint
byte //uint8
rune //int32
float32 float64
complex64 complex 128
*/

func main() {
	var mystring = "1"
	var x = 10
	var f float32 = 2.0
	fmt.Println(mystring,x,f)
	number,_ := strconv.Atoi(mystring)
	fmt.Println(number)
	x = number
	fmt.Println(x)
	fmt.PrintF("%T",f)
	fmt.PrintF("%v",f)
}
