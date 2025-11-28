package main

import "fmt"

func createTemperatureAdjuster() (func(adjustment float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(adjustment float64) float64 {
		return baseTemperature + adjustment
	}
	return adjustTemperature, baseTemperature
}

func main() {
	adjuster, baseTemp := createTemperatureAdjuster()
	fmt.Printf("Base Temperature: %.1f\n", baseTemp)
	fmt.Printf("Adjusted Temperature +1.5: %.1f grand C\n", adjuster(1.5))
	fmt.Printf("Adjusted Temperature -3.0: %.1f grand C\n", adjuster(-3.0))
	fmt.Printf("Adjusted Temperature 5.0: %.1f grand C\n", adjuster(5.0))
}
