package main

import (
	"fmt"
	"sort"
	
	"github.com/afharvey/ww-sweet-shop/packer"
)

func main() {
	kinds := []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000}
	//kinds := []int{250, 500, 1000, 2000, 5000}
	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i] > kinds[j] // descending
	})

	challenges := []int{1, 20, 66, 77, 100, 250, 251, 500, 501, 752, 1752, 5001, 12001, 100000, 100035, 100066}
	//challenges := []int{1, 250, 251, 501, 12001}
	for _, challenge := range challenges {

		fmt.Printf("challenge: %d   starting kinds: %v\n", challenge, kinds)
		for _, kind := range kinds {
			target := packer.MakeTarget(challenge, kind)
			fmt.Printf("%d:%d, ", kind, target)
		}
		orderSize, winnerKind := GetOrderSize(challenge, kinds)
		fmt.Printf("\nnew order: %d\t\tbest kind: %d\n", orderSize, winnerKind)
		newKinds := ExcludeOverSized(orderSize, winnerKind, kinds)
		fmt.Println("new kinds: ", newKinds)
		fmt.Println("to ship: ", Picker(orderSize, newKinds))
		fmt.Println()
		fmt.Println()
	}
}

// GetOrderSize returns the smallest order we can send to the customer. 
// No less than requested and no more than necessary.
func GetOrderSize(requested int, kinds []int) (int, int) {
	var orderSize int
	var winner int
	for _, i := range kinds {
		target := packer.MakeTarget(requested, i)
		if target < orderSize || orderSize == 0 {
			winner = i
			orderSize = target
		}
	}
	return orderSize, winner
}

// Exclude kind of boxes which are larger than we require.
func ExcludeOverSized(target, baseKind int, kinds []int) []int {
	o := make([]int, 0)
	for _, i := range kinds {

		if i > target {
			continue
		}

		if i == target {
			return []int{i} // quit early
		}

		// Exclude kinds which are not multiples of baseKind.
		if i != baseKind && i%baseKind != 0 {
			continue
		}

		o = append(o, i)
	}
	return o
}

func Picker(target int, kinds []int) []int {
	output := make([]int, 0)
	remaining := target

	for _, kind := range kinds {

		numberOfBoxes := remaining / kind
		remaining -= kind * numberOfBoxes

		for i := 0; i < numberOfBoxes; i++ {
			output = append(output, kind)
		}

	}

	return output
}
