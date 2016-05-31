package main

import (
	"fmt"
)

func main() {
	carLotA := 10
	carLotB := 20

	carLotTotal := carLotA + carLotB
	fmt.Println("carLotTotal +:",carLotTotal) 

	carLotTotal = carLotTotal - 10
	fmt.Println("carLotTotal -:",carLotTotal) 

	carLotTotal *= 10
	fmt.Println("carLotTotal *:",carLotTotal) 

	carLotTotal = carLotB / carLotA
	fmt.Println("carLotTotal /:",carLotTotal) 

	carLotTotal = carLotB % carLotA
	fmt.Println("carLotTotal %:",carLotTotal) 
}
