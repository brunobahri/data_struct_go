package main

import "fmt"

// Item represents an item with a value and a weight
type Item struct {
	value  int
	weight int
}

// knapsack function solves the 0/1 knapsack problem using dynamic programming
func knapsack(items []Item, capacity int) int {
	n := len(items)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if items[i-1].weight <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-items[i-1].weight]+items[i-1].value)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

// max function returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	items := []Item{
		{value: 60, weight: 10},
		{value: 100, weight: 20},
		{value: 120, weight: 30},
	}

	capacity := 50
	fmt.Printf("Maximum value in Knapsack = %d\n", knapsack(items, capacity))
}
