package main

import (
	"fmt"
)

// calculateCosts calculates and returns the structural and labor costs.
// areaSqm is the area in square meters, ratePerSqm is the rate per square meter in PHP,
// and laborPercentage is the percentage of the structural cost that will go towards labor.
func calculateCosts(areaSqm float64, ratePerSqm float64, laborPercentage float64) (float64, float64, float64) {
	structuralCost := areaSqm * ratePerSqm
	laborCost := structuralCost * (laborPercentage)
	totalCost := structuralCost + laborCost
	return structuralCost, laborCost, totalCost
}

func main() {
	areaSqm := 80.0         // Example area in square meters
	ratePerSqm := 12000.0   // Example rate per square meter in PHP
	laborPercentage := 0.45 // Labor cost as a percentage of the structural cost

	structuralCost, laborCost, totalCost := calculateCosts(areaSqm, ratePerSqm, laborPercentage)

	fmt.Printf("Structural Cost: Php %.2f\n", structuralCost)
	fmt.Printf("Labor Cost: Php %.2f\n", laborCost)
	fmt.Printf("Total Cost: Php %.2f\n", totalCost)
}
