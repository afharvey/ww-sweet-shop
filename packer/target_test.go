package packer_test

import (
	"github.com/afharvey/ww-sweet-shop/packer"
	"testing"
)

func TestMakeTarget(t *testing.T) {
	tests := []struct {
		name                string
		requested, unitSize int
		expected            int
	}{
		{
			name:      "1 -> 250",
			requested: 1,
			unitSize:  250,
			expected:  250,
		},
		{
			name:      "250 -> 250",
			requested: 250,
			unitSize:  250,
			expected:  250,
		},
		{
			name:      "251 -> 500",
			requested: 251,
			unitSize:  250,
			expected:  500,
		},
		{
			name:      "501 -> 750",
			requested: 501,
			unitSize:  250,
			expected:  750,
		},
		{
			name:      "12001 -> 12250",
			requested: 12001,
			unitSize:  250,
			expected:  12250,
		},
		{
			name:      "752 -> 1000",
			requested: 752,
			unitSize:  250,
			expected:  1000,
		},
		{
			name:      "1752 -> 2000",
			requested: 1752,
			unitSize:  250,
			expected:  2000,
		},
		{
			name:      "7777 -> 8000",
			requested: 7777,
			unitSize:  250,
			expected:  8000,
		},
		{
			name:      "1 -> 888",
			requested: 1,
			unitSize:  888,
			expected:  888,
		},
		{
			name:      "1, 888 -> 888",
			requested: 888,
			unitSize:  1,
			expected:  888,
		},
		{
			name:      "4, 888 -> 888",
			requested: 888,
			unitSize:  4,
			expected:  888,
		},
		{
			name:      "88, 888 -> 968",  // 888/88=10.09 ~ 11
			requested: 888,
			unitSize:  88,
			expected:  968,
		},
		{
			name:      "88 -> 888",
			requested: 88,
			unitSize:  888,
			expected:  888,
		},
		{
			name:      "250, 251 -> 500",
			requested: 251,
			unitSize:  250,
			expected:  500,
		},
		{
			name:      "10, 251 -> 260",
			requested: 251,
			unitSize:  10,
			expected:  260,
		},
		{
			name:      "5, 251 -> 255",
			requested: 251,
			unitSize:  5,
			expected:  255,
		},
		{
			name:      "2, 251 -> 252",
			requested: 251,
			unitSize:  2,
			expected:  252,
		},
		{
			name:      "1, 251 -> 251",
			requested: 251,
			unitSize:  1,
			expected:  251,
		},
		{
			name: "33, 12001 -> 12012",
			requested: 12001,
			unitSize: 33,
			expected: 12012,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := packer.MakeTarget(test.requested, test.unitSize); test.expected != result {
				t.Logf("wanted %d got %d", test.expected, result)
				t.FailNow()
			}
		})
	}
}
