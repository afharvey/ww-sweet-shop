package packer_test

import (
	"github.com/afharvey/ww-sweet-shop/packer"
	"github.com/go-test/deep"
	"reflect"
	"testing"
)

func TestMakeOrder(t *testing.T) {
	tests := []struct {
		name      string
		challenge int
		kinds     []int
		expected  []int
	}{
		// Examples from the spec
		{
			name:      "1: 1x250",
			challenge: 1,
			kinds:     []int{250, 500, 1000, 2000, 5000},
			expected:  []int{250},
		},
		{
			name:      "250: 1x250",
			challenge: 250,
			kinds:     []int{250, 500, 1000, 2000, 5000},
			expected:  []int{250},
		},
		{
			name:      "251: 1x500",
			challenge: 251,
			kinds:     []int{250, 500, 1000, 2000, 5000},
			expected:  []int{500},
		},
		{
			name:      "501: 1x500, 1x250",
			challenge: 501,
			kinds:     []int{250, 500, 1000, 2000, 5000},
			expected:  []int{500, 250},
		},
		{
			name:      "12001: 2x5000, 1x2000, 1x250",
			challenge: 12001,
			kinds:     []int{250, 500, 1000, 2000, 5000},
			expected:  []int{5000, 5000, 2000, 250},
		},
		// Interesting numbers
		{
			name:      "1: 1x33",
			challenge: 1,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{33},
		},
		{
			name:      "20: 1x33",
			challenge: 20,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{33},
		},
		{
			name:      "66: 1x66",
			challenge: 66,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{66},
		},
		{
			name:      "77: 1x99",
			challenge: 77,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{99},
		},
		{
			name:      "100: 3x34",
			challenge: 100,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{34, 34, 34},
		},
		{
			name:      "250: 1x250",
			challenge: 250,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{250},
		},
		{
			name:      "251: 4x66",
			challenge: 251,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{66, 66, 66, 66},
		},
		{
			name:      "500: 1x500",
			challenge: 500,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{500},
		},
		{
			name:      "501: 15x34",
			challenge: 501,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		},
		{
			name:      "752: 7x99,1x66",
			challenge: 752,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{99, 99, 99, 99, 99, 99, 99, 66},
		},
		{
			name:      "100000: 10x10000",
			challenge: 100000,
			kinds:     []int{33, 34, 66, 99, 250, 500, 1000, 2000, 5000, 6000, 7000, 8000, 9000, 10000},
			expected:  []int{10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := packer.MakeOrder(test.challenge, test.kinds)
			if !reflect.DeepEqual(test.expected, result) {
				t.Log(deep.Equal(test.expected, result))
				t.FailNow()
			}
		})
	}
}
