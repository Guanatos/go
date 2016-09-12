package main

import "fmt"

func main(){
	d := dealerShip{name:"Toyota",city:"Guanatos"}
	fmt.Println("name:",d.name)
	fmt.Println("city:",d.city)

	var d2 dealerShip
	d2 = dealerShip{name:"Honda",city:"Zapopan"}
	fmt.Println(d2)
}

type dealerShip struct{
	name string
	city string
}

