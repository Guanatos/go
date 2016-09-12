package main

import "fmt"

// greater than (>) or less than (<)

func main() {
	a := [4]int{4,3,2,1}
	fmt.Println("Elements: ",a)
	n := len(a)
	fmt.Println("Lenght (n): ",n)
	k := 3
	fmt.Println("Subarray Lenght (k): ",k)
	fmt.Println("########")
	for i := 0; i <= n-k; i++ {
	    fmt.Println("Value (i): ",i)
	    max := a[i]
	    fmt.Println("Element a[i]: ",a[i])
	    fmt.Println("Inital Max: ",max)
	    for j := 1; j < k; j++ {
	    	fmt.Println("  Value (j): ",j)
	    	fmt.Println("  Value (k): ",k)
	    	fmt.Println("  Value (j < k): ",j<k)
                if (a[i+j] > max){
			fmt.Println("Value (a[i+j] > max): ",a[i+j] > max)
			max = a[i+j]
	                fmt.Println("New Max: ",max)
		    }
		}
		fmt.Println("Final Max: ",max)
	}
}
