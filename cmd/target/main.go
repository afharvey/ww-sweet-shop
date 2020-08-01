package main

import (
	"fmt"
	"github.com/afharvey/ww-sweet-shop/packer"
	"sort"
)

func main() {
	kinds := []int{33, 34, 250, 500, 1000, 2000, 5000, 6000, 7000, 6000, 7000, 8000, 9000, 10000}
	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i] > kinds[j] // descending
	})
	smallest := kinds[len(kinds)-1]

	//challenges := []int{1, 20, 66, 77, 100, 250, 251, 500, 501, 752, 1752, 5001, 12001, 100000}
	challenges := []int{752}
	for _, challenge := range challenges {
		target := packer.MakeTarget(challenge, smallest)
		fmt.Printf("challenge: %d   old challenge: %d\n", challenge, target)
		MakeOrder(challenge, kinds)

		fmt.Println()
		fmt.Println()
	}
}

func MakeOrder(want int, kinds []int) {

	// preprocess - exclude kinds we dont need or quit early if we can.
	if QuitEarly(want, kinds) {
		return
	}
	k, quit := PruneKinds(want, kinds)
	if quit {
		return
	}

	a, b := Partition(k)

	fmt.Println("target: ", want)
	fmt.Println("kinds: ", k)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	var iterations int

outer:
	for _, aa := range a {
		for _, bb := range b {

			iterations++

			sum := aa + bb
			if sum == 0 || sum > want {
				continue
			}

			fmt.Printf("- \twant:%d\ta:%d\tb:%d\tsum:%d  \tdemo:%d\n", want, aa, bb, sum, aa+(bb*2))

			want = want - sum
			if want > kinds[len(kinds)-1] {
				goto outer
			}
		}
	}

	fmt.Printf("final:\t%d\t\treached in iterations:%d\n", want, iterations)
}

func QuitEarly(want int, kinds []int) bool {

	// preprocess - exclude kinds we dont need or quit early if we can.
	smallest := kinds[len(kinds)-1]
	if want <= smallest {
		fmt.Printf("quit early! want:%d found:%d\n", want, smallest)
		return true
	}
	return false
}

func PruneKinds(want int, kinds []int) ([]int, bool) {
	k := make([]int, 0)
	for _, i := range kinds {

		if want == i {
			fmt.Printf("quit early! want:%d found:%d\n", want, i)
			return nil, true
		}

		if i > want {
			continue
		}
		k = append(k, i)
	}
	return k, false
}

func Partition(k []int) ([]int, []int) {
	a := make([]int, len(k)+1) // appends zero
	b := make([]int, len(k)+1)
	copy(a, k)
	copy(b, k)
	return a, b
}
