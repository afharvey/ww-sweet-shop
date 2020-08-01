package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	want := 1752

	kinds := []int{250, 500, 1000, 2000, 5000}

	sort.Slice(kinds, func(i, j int) bool {
		return i > j
	})
	smallest := kinds[len(kinds)-1]
	largest := kinds[0]
	second := kinds[len(kinds)-2]

	a := make([]int, len(kinds)+1) // appends zero
	b := make([]int, len(kinds)+1)

	copy(a, kinds)
	copy(b, kinds)

	fmt.Println("want \t", want)
	fmt.Println("kinds\t", kinds)
	fmt.Printf("smallest:%d\tlargest:%d\tsecond:%d\n", smallest, largest,second)
	fmt.Println("a\t", a)
	fmt.Println("b\t", b)

next:
	best := &Pair{}
	for _, aa := range a {
		for _, bb := range b {

			candidate := &Pair{aa, bb}

			// Step 1, identify a candidate.
			// The sum of a and b is too much so ignore them.
			//if candidate.Sum() > want {
			//	continue
			//}

			// Step 2, identify the best candidate.
			// Is this sum the best candidate we know so far?
			if best.IsBetter(want, candidate) {
				continue
			}

			if want < second && want > smallest {
				want = want - second
				fmt.Printf("want:%d\ta:%d\tb:%d\tsecond:%d    \tok:%v\n", want, aa, bb, second, second <= want)
				break
			}

			fmt.Printf("want:%d\ta:%d\tb:%d\tsum:%d    \tok:%v\n", want, aa, bb, candidate.Sum(), candidate.Sum() <= want)
			want = want - candidate.Sum()
			if want > 0 {

				if want < smallest {
					fmt.Printf("tail:\t%d\tsmallest:\t%d\n", want, smallest)
					want = want - smallest
				}
				//if want < smallest {
				//	break
				//}

				time.Sleep(1 * time.Second)
				goto next
			}
		}
	}
	
	//if want > 0 {
	//	var leftOver int
	//	for _, i := range kinds {
	//		if i > want {
	//			leftOver = i
	//		}
	//	}
	//	want = want - leftOver
	//	fmt.Printf("left overs %d\n", leftOver)
	//}
	

	fmt.Printf("final:\t%d\n", want)
}

type Pair struct {
	a, b int
}

func (p *Pair) Sum() int {
	return p.a + p.b
}

func (p *Pair) IsBetter(target int, other *Pair) bool {
	if other.Sum() > target {
		return true
	}

	if p.Sum() > other.Sum() {
		return true
	}
	return false
}
