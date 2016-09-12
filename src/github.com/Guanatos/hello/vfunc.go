package main

import "fmt"

// Main function
func main(){
	d1 := "Toyota"
	d2 := "Honda"
	d3 := "Nissan"
	d4 := "Rover"

	printDealers(d1,d2,d3,d4)

	fmt.Println("\n")
	
	dealers := []string{d1,d2,d3,d4}
	printDealers(dealers...)
}

// Function to print dealers which is an array of strings
func printDealers(dealers...string){
	for _,dealerName := range dealers {
		fmt.Println(dealerName)
	}
}
