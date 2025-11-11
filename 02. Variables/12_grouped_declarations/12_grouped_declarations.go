package main

import "fmt"

func main() {
	var coffeeType = "Latte"
	var quantity = 2
	var unitPrice = 4.50

	fmt.Printf("Order: %d x %s at $%.2f each\n", quantity, coffeeType, unitPrice)

	var (
		customerName      = "Alice"
		tableNumber  int  = 5
		isReadyToPay bool = true
	)
	fmt.Printf("Customer: %s, Table: %d, Ready to Pay: %t\n", customerName, tableNumber, isReadyToPay)

	// no unused variable error for constants
	const (
		SizeSmall  = "S"
		SizeMedium = "M"
		SizeLarge  = "L"
	)
}
