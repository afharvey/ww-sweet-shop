package packer

import (
	"fmt"
	"sort"
)

type Order struct {
	Desired float64
	Size    float64
	Boxes   []*Box
}

func NewOrder(d float64) *Order {
	return &Order{
		Desired: d,
		Boxes:   make([]*Box, 0),
	}
}

func (o *Order) Add(b *Box) {
	fmt.Printf("add %f\n", b.Size)
	o.Boxes = append(o.Boxes, b)
	o.Size += b.Size
}

func (o *Order) Remaining() float64 {
	if r := o.Desired - o.Size; r > 0 {
		return r
	}
	return 0
}

// An Box is a kind of box of sweets.
type Box struct {
	Size float64
}

// Designs the optimal order given a desired quantity and list of boxes we can use.
func BuildOrder(desired float64, kindsOfBox []*Box) (order *Order) {
	order = NewOrder(desired)

	sort.Slice(kindsOfBox, func(i, j int) bool {
		return kindsOfBox[i].Size > kindsOfBox[j].Size // descending
	})

	for _, v := range kindsOfBox {
		box := v
		
		if order.Remaining() >= box.Size {
			j := int(order.Remaining()) / int(box.Size)
			for i:=0;i<j;i++{
				newBox := v
				order.Add(newBox)
			}
		}
	}

	return
}

//https://www.geeksforgeeks.org/find-number-currency-notes-sum-upto-given-amount/
//https://www.w3resource.com/c-programming-exercises/basic-declarations-and-expressions/c-programming-basic-exercises-16.php
