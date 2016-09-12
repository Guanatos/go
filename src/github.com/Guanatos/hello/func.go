package main

import "fmt"

// Main body
func main(){
	d1 := dealerShip{"Ford","La Barca"}
	dealerName := createDealerFullName(d1.name,d1.city)
	fmt.Println(dealerName)

	sold,name := calcInvUtil(100,200,d1)
	fmt.Println(name,sold)
}

// Strutures
type dealerShip struct{
	name string
	city string
}

// Functions
func createDealerFullName(s1 string, s2 string) string {
     return s1 + " of " +s2
}

func calcInvUtil(remaining float32, start float32, dealer dealerShip) (percSold float32, dealerName string){
	dealerName = dealer.name+" sold "
	percSold = 1 - remaining/start
	return
}
