package main

import "fmt"

func main() {
	coffee := "Cappuccino"
	pointer := &coffee

	fmt.Println("Coffee name:", coffee)
	fmt.Println("Memory address of coffee variable:", pointer)
	fmt.Printf("Pointer address: %p\n", pointer)

	fmt.Println("--------------------------------------------------")
	coffeeCopy := coffee // create a copy of the value

	fmt.Println("Coffee name:", coffeeCopy)
	fmt.Println("Memory address of coffee variable:", &coffeeCopy)
	fmt.Printf("Pointer address: %p\n", &coffeeCopy)
}
