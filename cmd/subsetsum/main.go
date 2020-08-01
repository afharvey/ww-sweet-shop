package main

import "fmt"

func main() {
	
	SubsetSum(5, func(ints []int) {
		fmt.Println(ints)
	})
	
}

func SubsetSum(n int, emit func([]int)) {
	var subsetSum func([]int, int, int)
	subsetSum = func(a []int, i, sum int) {
		a[i] = sum
		emit(a[:i+1])
		for a[i]--; a[i] > 0; a[i]-- {
			subsetSum(a, i+1, sum-a[i])
		}
	}
	subsetSum(make([]int, n), 0, n)
}