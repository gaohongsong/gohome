package utils

import "testing"

type test struct {
	in  []float64
	out float64
}

var tests = []test{
	{[]float64{1, 2, 3, 4, 5}, 3},
	{[]float64{10, 12, 13, 5}, 10},
	{[]float64{10, 20, 30, 40, 50}, 30},
}

func TestAverate(t *testing.T) {
	for _, tc := range tests {
		if v := Average(tc.in...); v != tc.out {
			//t.Fatalf("Avg(%v), Expect %v, go %v", tc.in, tc.out, v)
			t.Errorf("Avg(%v), Expect %v, go %v", tc.in, tc.out, v)
		}
	}
}
