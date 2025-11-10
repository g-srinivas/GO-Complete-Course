package main

import "fmt"

func main() {
	price := 4.50 // float64

	quantity := 15 // int

	// total := price * quantity // ❌ Go does not convert types

	total := price * float64(quantity) // ✅ Explicit conversion

	fmt.Printf("Total Price: %.2f", total)
}
