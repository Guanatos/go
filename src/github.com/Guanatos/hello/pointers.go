package main

import "fmt"

var shippingDays = 30
var shippingDaysPtr = new(int)

func main(){
     shippingDaysAdjust(shippingDays)
     fmt.Println("after shipping Days Adjustments call",shippingDays)

     shippingDaysAdjustPtr(&shippingDays)
     fmt.Println("after shipping Days Adjustments Ptr call",shippingDays)
     fmt.Println("after shipping Days Adjustments Ptr call",&shippingDays)

     shipper := shipper{}
     shipper.name = "Datsun"
     shipper.id = 400

     fmt.Println("Shippers: ",shipper)

     shipperUpdates(&shipper)

     fmt.Println("Shippers Id: ",shipper.id)
     fmt.Println("Shippers Name: ",shipper.name)

//   Pointers
     *shippingDaysPtr = 55

//   We are not using & as it's already a pointer
     shippingDaysAdjustPtr(shippingDaysPtr)
     fmt.Println("shippingDaysPtr after",*shippingDaysPtr)
}

// Copy of the variable
func shippingDaysAdjust(shippingDays int){
     shippingDays += 10
}

// Pointer
func shippingDaysAdjustPtr(shippingDays *int){
     *shippingDays += 10
}

// Pointer to an structure
func shipperUpdates(s *shipper){
     s.id += 10
     s.name += " Inc."
}

type shipper struct{
     name string
     id int
}
