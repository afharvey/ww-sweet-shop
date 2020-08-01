package main

import "fmt"

func main() {
	
	want := 50001
	
	kinds := []int{250, 500, 1000, 2000, 5000}
	split := len(kinds[len(kinds)/2:])+1
	
	a := make([]int, split) 
	b := make([]int, split)
	
	copy(a, kinds[len(kinds)/2:])  // partition the set {1...n} into two sets A and B
	copy(b, kinds[:len(kinds)/2])
	
	fmt.Println("want \t", want)
	fmt.Println("kinds\t", kinds)
	fmt.Println("a\t", a)
	fmt.Println("b\t", b)
	
	for _, aa := range a {
		for _, bb := range b {
			sum := aa+bb
			if sum == 0 {
				continue
			}
			
			ok := sum <= want
			if !ok {
				continue
			}
			fmt.Printf("want:%d\ta:%d\tb:%d\tsum:%d    \tok:%v\n", want, aa, bb, sum, ok)
		}
	}
}
