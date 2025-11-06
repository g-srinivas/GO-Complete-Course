package main

import "fmt"

func main() {
	// Declaring a variable without initialization
	var price float64

	// Initializing the variable later
	price = 3.50
	fmt.Printf("Price of the coffee is: %.2f\n", price)
}
