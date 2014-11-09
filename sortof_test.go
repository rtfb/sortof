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

func TestWalkTree(t *testing.T) {
	var cases = []struct {
		input    *Tree
		expected string
	}{
		{nil, ""},
		{&Tree{"a", nil}, "a"},
		{&Tree{"a", []*Tree{&Tree{"b", nil}}}, "a\n.\tb"},
		{&Tree{"a", []*Tree{&Tree{"b", nil}, &Tree{"c", nil}}}, "a\n.\tb\n.\tc"},
		{&Tree{"a", []*Tree{&Tree{"b",
			[]*Tree{&Tree{"x", nil}, &Tree{"y", nil}, &Tree{"z", nil}}},
			&Tree{"c", nil}}}, "a\n.\tb\n.\t.\tx\n.\t.\ty\n.\t.\tz\n.\tc"},
	}
	for _, c := range cases {
		result := printTree(c.input)
		if result != c.expected {
			t.Fatalf("Bad result walking tree %#v. Expected\n%v\n, but got\n%v",
				c.input, c.expected, result)
		}
	}
}

func TestCompareTrees(t *testing.T) {
	var cases = []struct {
		t1       *Tree
		t2       *Tree
		expected bool
	}{
		{nil, nil, true},
		{&Tree{"a", nil}, nil, false},
		{nil, &Tree{"a", nil}, false},
		{&Tree{"a", nil}, &Tree{"a", nil}, true},
		{&Tree{"a", nil}, &Tree{"b", nil}, false},
		{&Tree{"a", []*Tree{&Tree{"b", nil}}},
			&Tree{"a", []*Tree{&Tree{"b", nil}}}, true},
		{&Tree{"a", []*Tree{&Tree{"b", nil}}},
			&Tree{"x", []*Tree{&Tree{"b", nil}}}, false},
		{&Tree{"a", []*Tree{&Tree{"b", nil}, &Tree{"c", nil}}},
			&Tree{"a", []*Tree{&Tree{"b", nil}, &Tree{"c", nil}}}, true},
	}
	for _, c := range cases {
		result := compareTrees(c.t1, c.t2)
		if result != c.expected {
			t.Fatalf("Wrong result comparing trees %#v and %#v. Expected %v, but got %v",
				c.t1, c.t2, c.expected, result)
		}
	}
}

func TestParseStringToTree(t *testing.T) {
	var cases = []struct {
		input    string
		expected *Tree
	}{
		{"", nil},
		{"()", nil},
		{"(a)", &Tree{"a", nil}},
		{"(a(b))", &Tree{"a", []*Tree{&Tree{"b", nil}}}},
		{"(a(b c))", &Tree{"a", []*Tree{&Tree{"b", nil}, &Tree{"c", nil}}}},
		{"(a(b(x y)c))", &Tree{"a", []*Tree{&Tree{"b",
			[]*Tree{&Tree{"x", nil}, &Tree{"y", nil}}},
			&Tree{"c", nil}}}},
		{"(a(b(x y)c d))", &Tree{"a", []*Tree{&Tree{"b",
			[]*Tree{&Tree{"x", nil}, &Tree{"y", nil}}},
			&Tree{"c", nil}, &Tree{"d", nil}}}},
		{"(a(b(x y)c d(q)))", &Tree{"a", []*Tree{&Tree{"b",
			[]*Tree{&Tree{"x", nil}, &Tree{"y", nil}}},
			&Tree{"c", nil}, &Tree{"d", []*Tree{&Tree{"q", nil}}}}}},
	}
	for _, c := range cases {
		result := parseTree(c.input)
		if !compareTrees(result, c.expected) {
			t.Fatalf("Wrong result parsing tree %q. Expected %v, but got %v",
				c.input, printTree(c.expected), printTree(result))
		}
	}
}
