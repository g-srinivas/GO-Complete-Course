package main

import "fmt"

func main() {
	// Explicitly declaring variable type and assigning value
	var coffee_name string = "Espresso"

	// Type inference - Go infers the type based on the assigned value
	var coffee_price = 2.50

	// Using short variable declaration for size. This can only be used inside functions.
	size := "Medium"
	fmt.Println(size, coffee_name, "price is $", coffee_price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffee_name, coffee_price)
}
