package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	kinds := []int{250, 500, 1000, 2000, 5000}
	sort.Slice(kinds, func(i, j int) bool {
		return i > j
	})

	challenges := []int{1, 250, 251, 500, 501, 752, 1752, 12001}
	for _, challenge := range challenges {
		MakeOrder(challenge, kinds)
	}
}

func MakeOrder(want int, kinds []int) {

	a := make([]int, len(kinds)+1) // appends zero
	b := make([]int, len(kinds)+1)

	copy(a, kinds)
	copy(b, kinds)

	fmt.Println()
	fmt.Println()
	fmt.Println("want \t", want)
	fmt.Println("kinds\t", kinds)
	fmt.Println("a\t", a)
	fmt.Println("b\t", b)
	
outer:
	best := &Pair{}
	for _, aa := range a {
		for _, bb := range b {

			candidate := &Pair{aa, bb}
			if candidate.Sum()==0{
				continue
			}
			
			if best.IsBetter(want, candidate, candidate.Sum()==kinds[len(kinds)-1]) {
				continue
			}

			fmt.Printf("want:%d   \ta:%d\tb:%d\tsum:%d    \tok:%v\n", want, aa, bb, candidate.Sum(), candidate.Sum() <= want)
			want = want - candidate.Sum()
			if want > kinds[len(kinds)-1] {
				time.Sleep(1 * time.Second)
				goto outer
			}
		}
	}
	
	fmt.Printf("final:\t%d\n", want)
}

type Pair struct {
	a, b int
}

func (p *Pair) Sum() int {
	return p.a + p.b
}

func (p *Pair) IsBetter(target int, other *Pair, last bool) (better bool) {

	if other.Sum() > target || p.Sum() > other.Sum() {
		better = true
	}
	
	fmt.Printf("\t\tIsBetter? %s vs %s \t%v\t%v\n", p, other, last, better)
	return better
}

func (p *Pair) String() string {
	return fmt.Sprintf("%d, %d", p.a, p.b)
}