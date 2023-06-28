package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define the possible symbols and their weights
	symbols := []string{"apple", "banana", "strawberry", "honeydew", "watermelon", "pineapple"}
	weights := []int{1, 10, 50, 100, 200, 800}

	// Initialize the slot machine with a random generator
	rand.Seed(time.Now().UnixNano())

	// Play the slot machine 1000 times and count the wins
	wins := make(map[string]int)
	for i := 0; i < 1000; i++ {
		// Spin the slot machine
		result := spin(symbols, weights)

		// Count the wins
		for _, symbol := range result {
			wins[symbol]++
		}
	}

	// Print the probabilities of winning for each symbol
	for _, symbol := range symbols {
		fmt.Printf("%s: %.2f%%\n", symbol, float64(wins[symbol])/10.0)
	}
}

// spin simulates a single spin of the slot machine
func spin(symbols []string, weights []int) []string {
	// Select a random symbol according to the weights
	symbol := symbols[weightedRand(weights)]

	// Return the result of the spin
	return []string{symbol, symbol, symbol}
}

// weightedRand returns a random index according to the given weights
func weightedRand(weights []int) int {
	// Calculate the total weight
	total := 0
	for _, weight := range weights {
		total += weight
	}

	// Select a random value between 0 and the total weight
	value := rand.Intn(total)

	// Find the corresponding index
	for i, weight := range weights {
		if value < weight {
			return i
		}
		value -= weight
	}
	return len(weights) - 1
}
