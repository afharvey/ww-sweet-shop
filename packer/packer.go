package packer

import (
	"fmt"
	"sort"
)

type Order struct {
	Desired int
	Size    int
	Boxes   []*Box
}

func NewOrder(d int) *Order {
	return &Order{
		Desired: d,
		Boxes:   make([]*Box, 0),
	}
}

func (o *Order) Add(b *Box) {
	fmt.Printf("add %d\n", b.Size)
	o.Boxes = append(o.Boxes, b)
	o.Size += b.Size
}

// An Box is a kind of box of sweets.
type Box struct {
	Size int
}

// Designs the optimal order given a desired quantity and list of boxes we can use.
func BuildOrder(desired int, kindsOfBox []*Box) (order *Order) {
	order = NewOrder(desired)

	sort.Slice(kindsOfBox, func(i, j int) bool {
		return kindsOfBox[i].Size > kindsOfBox[j].Size // descending
	})

	for _, v := range kindsOfBox {
		box := v

		remaining := order.Desired - order.Size
		if remaining < 1 {
			return
		}

		if remaining == box.Size {
			order.Add(box)
			return
		}

		// How many of this box can we fit into the order?
		boxesToAdd := remaining / box.Size
		fmt.Printf("desired:%d remaining:%d size:%d, want to add %d boxes\n", order.Desired, remaining, box.Size, boxesToAdd)
		if boxesToAdd > 0 {
			for i := 0; i < boxesToAdd; i++ {
				order.Add(box)
			}
			continue
		}
		
		previous := kindsOfBox[len(kindsOfBox)-1]
		order.Add(previous)
		return 
	}

	fmt.Println("out")

	// No more boxes would fit so the best we can do is add the smallest.
	smallest := kindsOfBox[len(kindsOfBox)-1]
	order.Add(smallest)
	return
}
