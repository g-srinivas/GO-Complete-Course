package main

import "fmt"

func main() {
	var totalOrders int     // default value is 0
	var customerName string // default value is ""
	var isOrderReady bool   // default value is false
	fmt.Printf("Total Orders: %d\n", totalOrders)
	fmt.Printf("Customer Name: '%s'\n", customerName)
	fmt.Printf("Is Order Ready: %t\n", isOrderReady)
}
