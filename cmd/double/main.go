package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	want := 251

	kinds := []int{250, 500, 1000, 2000, 5000}

	sort.Slice(kinds, func(i, j int) bool {
		return i > j
	})
	smallest := kinds[len(kinds)-1]

	a := make([]int, len(kinds)+1) // appends zero
	b := make([]int, len(kinds)+1)

	copy(a, kinds)
	copy(b, kinds)

	fmt.Println("want \t", want)
	fmt.Println("kinds\t", kinds)
	fmt.Println("a\t", a)
	fmt.Println("b\t", b)

next:
	best := &Pair{} 
	for _, aa := range a {
		for _, bb := range b {
			sum := aa + bb

			// Step 1, identify a candidate.
			// The sum of a and b is too much so ignore them.
			if sum > want {
				continue
			}

			// Step 2, identify the best candidate.
			// Is this sum the best candidate we know so far?
			if best.Sum() > sum {
				continue
			}
			
			best = &Pair{aa,bb}
			
			fmt.Printf("want:%d\ta:%d\tb:%d\tsum:%d    \tok:%v\n", want, aa, bb, sum, sum <= want)
			want = want - sum
			if want > 0 {
				
				if want < smallest {
					fmt.Printf("tail:\t%d\tsmallest:\t%d\n", want, smallest)
					want = want - smallest
				}
				
				time.Sleep(1 * time.Second)
				goto next
			}
		}
	}

	fmt.Printf("after loops %d\n", want)


	fmt.Printf("final:\t%d\n", want)
}

type Pair struct {
	a, b int
}

func (p *Pair) Sum() int {
	return p.a + p.b
}

