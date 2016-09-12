package main

import "fmt"
import "time"
<<<<<<< HEAD

func main() {
	fmt.Printf("Hello, Daniel\n")
	fmt.Printf("The time is",time.Now())
=======
import "runtime"

func main() {
	fmt.Println("Hello, Daniel\n")
	fmt.Printf("The time is:", time.Now())
	fmt.Println("\n")
	fmt.Printf("The CPU is:", runtime.NumCPU())
	fmt.Println("\n")
	fmt.Printf("The MEM is:", runtime.GOMAXPROCS(runtime.NumCPU()))
	fmt.Println("\n")
>>>>>>> 26dbb0736e0491f84ad172ef191d3274669052f9
}
