package packer_test

import (
	"github.com/afharvey/ww-sweet-shop/packer"
	"github.com/go-test/deep"
	"reflect"
	"testing"
)

// Example table from the brief
var exampleKinds = []*packer.Box{
	{Size: 250},
	{Size: 500},
	{Size: 1000},
	{Size: 2000},
	{Size: 5000},
}

func TestBuildOrder(t *testing.T) {
	tests := []struct {
		Name     string
		Kinds    []*packer.Box
		Quantity int
		Expected *packer.Order
	}{
		//{
		//	Name:     "1, first example where we immediately overflow",
		//	Kinds:    exampleKinds,
		//	Quantity: 1,
		//	Expected: &packer.Order{Desired: 1, Size: 250, Boxes: []*packer.Box{{Size: 250}}},
		//},
		//{
		//	Name:     "250, second example where we can use one box",
		//	Kinds:    exampleKinds,
		//	Quantity: 250,
		//	Expected: &packer.Order{Desired: 250, Size: 250, Boxes: []*packer.Box{{Size: 250}}},
		//},
		{
			Name:     "251, third example where we can use one box but must find it",
			Kinds:    exampleKinds,
			Quantity: 251,
			Expected: &packer.Order{Desired: 251, Size: 500, Boxes: []*packer.Box{{Size: 500}}},
		},
		//{
		//	Name:     "501, fourth example where we must combine boxes",
		//	Kinds:    exampleKinds,
		//	Quantity: 501,
		//	Expected: &packer.Order{Desired: 501, Size: 750, Boxes: []*packer.Box{{Size: 500}, {Size: 250}}},
		//},
		//{
		//	Name:     "12001, fifth example where we must combine more boxes",
		//	Kinds:    exampleKinds,
		//	Quantity: 12001,
		//	Expected: &packer.Order{Desired: 12001, Size: 12250, Boxes: []*packer.Box{
		//		{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250},
		//	}},
		//},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := packer.BuildOrder(test.Quantity, test.Kinds)
			if !reflect.DeepEqual(test.Expected, result) {
				t.Log(deep.Equal(test.Expected, result))
				t.FailNow()
			}
		})
	}
}
