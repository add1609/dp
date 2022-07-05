package fib

import "testing"

func TestFib(t *testing.T) {
	for _, test := range []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}{
		got := fib(test.n)
		if got != test.want {
			t.Errorf(
				"error:\ngot: %v\nwant: %v\n",
				got, test.want,
			)
		}
	}
}
