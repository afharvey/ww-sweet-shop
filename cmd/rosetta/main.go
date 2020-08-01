package main

import "fmt"

func main() {
	amount := 251
	fmt.Println("amount, ways to make change:", amount, countChange(amount))
}

func countChange(amount int) int64 {
	ways := make([]int64, amount+1)
	ways[0] = 1
	for _, coin := range []int{1000, 500, 250} {
		for j := coin; j <= amount; j++ {
			ways[j] += ways[j-coin]
		}
	}
	return ways[amount]
}