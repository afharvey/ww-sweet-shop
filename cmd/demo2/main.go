package main

import (
	"fmt"
	"github.com/afharvey/ww-sweet-shop/packer"
	"github.com/go-test/deep"
	"reflect"
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
	
	// ordered, correct
	check := map[int]*packer.Order{
		1: {
			Desired: 1,
			Size: 250,
			Boxes: []*packer.Box{{Size: 250}},
		},
		250: {
			Desired: 250,
			Size: 250,
			Boxes: []*packer.Box{{Size: 250}},
		},
		251: {
			Desired: 251,
			Size: 500,
			Boxes: []*packer.Box{{Size: 500}},
		},
		//501: {
		//	Desired: 1,
		//	Size: 250,
		//	Boxes: []*packer.Box{
		//		{Size: 250}, {Size: 500},
		//	},
		//},
		//12001: {
		//	Desired: 1,
		//	Size: 250,
		//	Boxes: []*packer.Box{
		//		{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250},
		//	},
		//},
	}
	
	for ordered, correct := range check {
		solution := MakeOrder(ordered, kinds)
		if !reflect.DeepEqual(correct, solution) {
			fmt.Println(deep.Equal(correct,solution))
		}
	}
}

func MakeOrder(size int, boxes []*packer.Box) *packer.Order {
	o := packer.NewOrder(size)
	
	for _, b := range boxes {
		for b.Size <= o.Remaining() {
			o.Add(b.Copy())
		}
	}
	
	if o.Remaining() > 0 {
		o.Add(boxes[len(boxes)-1].Copy())
	}
	
	return o
}