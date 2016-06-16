package main

import "fmt"

func main() {
	a := [4]int{4,1,3,2}
	fmt.Println("Elements: ",a)
	n := len(a)
	fmt.Println("Lenght (n): ",n)
	k := 3
	fmt.Println("Subarray Lenght (k): ",k)
	for i := 0; i <= n-k; i++ {
	    fmt.Println("Value (i): ",i)
	    fmt.Println("Element: ",a[i])
	    max := a[i]
	    fmt.Println("Inital Max: ",max)
	    for j := 1; j < k; j++ {
	    	fmt.Println("Value (j < k): ",j<k)
	    	fmt.Println("Value (j): ",j)
                if (a[i+j] > max){
			fmt.Println("Value (a[i+j] > max:",a[i+j] > max)
			max = a[i+j]
	                fmt.Println("New Max: ",max)
		    }
		}
		fmt.Println("Final Max: ",max)
	}
}
