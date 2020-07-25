package main

import "github.com/afharvey/ww-sweet-shop/packer"

func main() {

	kinds := []*packer.Box{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}

	packer.BuildOrder(100, kinds)
}
