package main

import "fmt"
import "time"
import "runtime"

func main() {
	fmt.Println("Hello, Daniel\n")
	fmt.Printf("The time is:", time.Now())
	fmt.Println("\n")
	fmt.Printf("The CPU is:", runtime.NumCPU())
	fmt.Println("\n")
	fmt.Printf("The MEM is:", runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println("\n")
}
