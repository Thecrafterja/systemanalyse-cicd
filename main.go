package main

import (
	"cicd-example/calculator"
	"fmt"
)

func main() {
	println("--- Example CI/CD application ---")

	fmt.Println("10 + 11 = ", calculator.Add(10, 11))
	fmt.Println("100 - 99 = ", calculator.Subtract(100, 99))

	println("---------------------------------")
}
