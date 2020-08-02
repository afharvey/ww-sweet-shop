package main

import (
	"fmt"
	"github.com/afharvey/ww-sweet-shop/packer"
	"sort"
)

func main() {
	
	fmt.Printf(" === Examples from the spec === \n\n")
	kinds := []int{250, 500, 1000, 2000, 5000}
	challenges := []int{1, 250, 251, 501, 12001}
	for _, challenge := range challenges {
		runDemo(challenge, kinds)
	}
	
	fmt.Printf("\n\n === Interesting numbers === \n\n")
	kinds = []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000}
	challenges = []int{1, 20, 66, 77, 100, 250, 251, 500, 501, 752, 1752, 5001, 12001, 100000, 100035, 100066}
	for _, challenge := range challenges {
		runDemo(challenge, kinds)
	}
}

func runDemo(challenge int, kinds []int) {
	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i] > kinds[j] // descending
	})
	fmt.Printf("challenge: %d   starting kinds: %v\n", challenge, kinds)
	for _, kind := range kinds {
		target := packer.MakeTarget(challenge, kind)
		fmt.Printf("%d:%d, ", kind, target)
	}
	orderSize, winnerKind := packer.GetOrderSize(challenge, kinds)
	fmt.Printf("\nnew order: %d\t\tbest kind: %d\n", orderSize, winnerKind)
	newKinds := packer.ExcludeUnhelpfulKinds(orderSize, winnerKind, kinds)
	
	fmt.Println("new kinds: ", newKinds)
	fmt.Println("to ship: ", packer.MakeOrder(orderSize, newKinds))
	fmt.Println()
	fmt.Println()
}