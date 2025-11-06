package main

import "fmt"

func main() {
	// Explicitly declaring variable type
	var cupsQty int = 5

	// cupsQty = "six" // ❌ Compile-time error type int
	fmt.Println("Number of cups:", cupsQty)

	var wasProcessed = true

	// wasProcessed = "yes" // ❌ Compile-time error type bool
	fmt.Println("Was the order processed?", wasProcessed)
}
