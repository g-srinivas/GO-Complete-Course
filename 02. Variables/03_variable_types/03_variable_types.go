package main

import "fmt"

func main() {
	name := "Americano"
	price := 2.99 // default type is float64
	ready := true
	count := 5
	var stockCount int64 = 150

	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of price: %T\n", price)
	fmt.Printf("Type of ready: %T\n", ready)
	fmt.Printf("Type of count: %T\n", count)
	fmt.Printf("Type of stockCount: %T\n", stockCount)
}
