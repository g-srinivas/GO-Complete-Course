package main

import "fmt"

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func main() {
	var points = calculateLoyaltyPoints(100.0)
	fmt.Println("Loyalty Points Earned:", points)
}
