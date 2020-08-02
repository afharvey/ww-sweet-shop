package main

import (
	"fmt"
	"github.com/afharvey/ww-sweet-shop/packer"
	"sort"
)

func main() {
	kinds := []int{2, 33, 34, 250, 500, 1000, 2000, 5000, 6000, 7000, 6000, 7000, 8000, 9000, 10000}
	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i] > kinds[j] // descending
	})
	//smallest := kinds[len(kinds)-1]

	challenges := []int{1, 20, 66, 77, 100, 250, 251, 500, 501, 752, 1752, 5001, 12001, 100000, 100035}
	//challenges := []int{66, 1752}
	for _, challenge := range challenges {
		
		fmt.Printf("challenge: %d   starting kinds: %v\n", challenge, kinds)
		for _, kind := range kinds {
			target := packer.MakeTarget(challenge, kind)
			fmt.Printf("%d:%d, ", kind, target)
		}
		fmt.Println()
		
		//MakeOrder(challenge, kinds)
		//fmt.Println()
		fmt.Println()
	}
}

func MakeOrder(want int, kinds []int) {

	ordered := want

	smallest := kinds[len(kinds)-1]
	var totalSent int

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

	var iterations int

outer:
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	for _, aa := range a {
		for _, bb := range b {

			iterations++

			sum := aa + bb
			if sum == 0 || sum > want {
				continue
			}

			fmt.Printf("- \twant:%d\ta:%d\tb:%d\tsum:%d\n", want, aa, bb, sum)

			want = want - sum
			totalSent += sum
			if want > 0 {
				a, _ = PruneKinds(want, a)
				b, _ = PruneKinds(want, b)
				goto outer
			}
		}
	}

	if want > 0 {
		want = want - smallest
		totalSent += smallest
		fmt.Printf("smallest:\t%d\n", smallest)
	}

	fmt.Printf("orderd:\t%d\tsent:\t%d\tbudget:%d\t\treached in iterations:%d\n", ordered, totalSent, want, iterations)
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

	// Are there any kinds which divide without a remainder?
	l := make([]int, 0)
	for _, i := range kinds {
		if i == 0 {
			continue
		}

		x := want % i
		if x == 0 {
			l = append(l, i)
		}

	}
	if len(l) > 0 {
		l = append(l, 0)
		return l, false
	}

	// Otherwise continue with removing kinds which are too large to be useful.
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
