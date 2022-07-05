package stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	for _, test := range []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	} {
		got := climbStairs(test.n)
		if got != test.want {
			t.Errorf(
				"error:\ngot: %v\nwant: %v\n",
				got, test.want,
			)
		}
	}
}
