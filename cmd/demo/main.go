package main

import (
	"fmt"
	"github.com/afharvey/ww-sweet-shop/packer"
	"sort"
)

func main(){
	
	kinds := []*packer.Box{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}

	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i].Size > kinds[j].Size // descending
	})
	
	smallest := kinds[len(kinds)-1]
	
	for _, desired := range []int{1, 250, 251, 300, 499, 501, 12001} {

		order := packer.NewOrder(desired)

		fmt.Printf("des\tbox\ttotal\tremain\tfits?\tquotient i f\tmod\n")

		for _, box := range kinds {

			fits := order.Remaining()-box.Size >= 1

			quotientI := order.Remaining() / box.Size // integer
			quotientF := float64(order.Remaining()) / float64(box.Size) // float
			modulo := order.Remaining() % box.Size

			fmt.Printf("%d\t%d\t%d\t%d\t%v\t%d %.2f    \t%d", desired, box.Size, order.Size, order.Remaining(), fits, quotientI, quotientF, modulo)

			for i := 0; i < quotientI; i++ {
				b := box
				order.Add(b)
			}

			fmt.Println()
		}
		
		if order.Remaining() > 0 {
			fits := order.Remaining()-smallest.Size >= 1
			quotientI := order.Remaining() / smallest.Size // integer
			quotientF := float64(order.Remaining()) / float64(smallest.Size) // float
			modulo := order.Remaining() / smallest.Size
			fmt.Printf("%d\t%d\t%d\t%d\t%v\t%d %.2f    \t%d\toverflow", desired, smallest.Size, order.Size, order.Remaining(), fits, quotientI, quotientF, modulo)
			order.Add(smallest)
		}

		
		fmt.Printf("\n----------\n")
		fmt.Printf("Final State\twant:%d\t got:%d\tdiff:%d\tboxes:%d", desired, order.Size, order.Desired-order.Size, len(order.Boxes))
		fmt.Println()
		fmt.Println()
	}
}