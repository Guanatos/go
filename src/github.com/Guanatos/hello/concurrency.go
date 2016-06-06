package main

import "fmt"

func main(){
//	foo()
//	bar()
	go foo2()
	go bar2()
}

func foo(){
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:",i)
	}
}

func bar(){
        for i := 0; i < 45; i++ {
                fmt.Println("Bar:",i)
        }
}

func foo2(){
        for i := 0; i < 45; i++ {
                fmt.Println("Foo:",i)
        }
}

func bar2(){
        for i := 0; i < 45; i++ {
                fmt.Println("Foo:",i)
        }
}
