package trib

import "testing"

func TestTrib(t *testing.T) {
	for _, test := range []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 7},
		{6, 13},
		{7, 24},
		{8, 44},
		{9, 81},
		{10, 149},
	}{
		got := trib(test.n)
		if got != test.want {
			t.Errorf(
				"error:\ngot: %v\nwant: %v\n",
				got, test.want,
			)
		}
	}
}
