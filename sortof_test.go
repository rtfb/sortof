package sortof

import (
	"fmt"
	"testing"
)

func TestInts(t *testing.T) {
	var cases = []struct {
		arr      []int
		expected string
	}{
		{arr: []int{}, expected: "[]"},
		{arr: []int{1}, expected: "[1]"},
		{arr: []int{2, 1}, expected: "[1 2]"},
		{arr: []int{1, 2}, expected: "[1 2]"},
		{arr: []int{1, 2, 3}, expected: "[1 2 3]"},
		{arr: []int{2, 1, 3}, expected: "[1 2 3]"},
		{arr: []int{3, 2, 1}, expected: "[1 2 3]"},
		{arr: []int{5, 6, 333, 8}, expected: "[5 6 8 333]"},
		{arr: []int{123, 124, 125, 126}, expected: "[123 124 125 126]"},
		{arr: []int{866, 700, 148, 587, 434, 898, 828, 893, 126, 657, 801, 868, 542},
			expected: "[126 148 434 542 587 657 700 801 828 866 868 893 898]"},
	}
	for i, c := range cases {
		Ints(c.arr)
		got := fmt.Sprintf("%v", c.arr)
		if got != c.expected {
			t.Fatalf("Test case %d failed.\nExpected: %v\nGot: %v\n",
				i, c.expected, got)
		}
	}
}
