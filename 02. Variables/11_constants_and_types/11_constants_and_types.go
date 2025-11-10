package main

import "fmt"

func main() {
	// Untyped constant with integer value
	const rewardPoints = 10
	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints)
	var totalRewardPoints float64 = 150.3
	// Adding untyped constant to float64 variable - valid, const adapts
	totalRewardPoints = totalRewardPoints + rewardPoints
}
