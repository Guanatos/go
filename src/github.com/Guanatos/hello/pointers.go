package main

import "fmt"

var shippingDays = 30

func main(){
     shippingDaysAdjust(shippingDays)
     fmt.Println("after shipping Days Adjustments call",shippingDays)

     shippingDaysAdjustPtr(&shippingDays)
     fmt.Println("after shipping Days Adjustments Ptr call",shippingDays)
     fmt.Println("after shipping Days Adjustments Ptr call",&shippingDays)

}

func shippingDaysAdjust(shippingDays int){
     shippingDays += 10
}

func shippingDaysAdjustPtr(shippingDays *int){
     *shippingDays += 10
}
