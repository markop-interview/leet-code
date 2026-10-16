package containsduplicate

import (
	"testing"
)

/*
Write table-driven tests with t.Run subtests covering at least:
- empty slice,
- single element,
- all distinct,
- duplicate at the start,
- duplicate at the end,
- negative numbers,
- large values near the constraints.

Both implementations should pass the same test table.
*/

var tests = []struct {
	name string
	s    []int
	want bool
}{
	{
		name: "empty slice",
		s:    []int{},
		want: false,
	},
	{
		name: "single element",
		s:    []int{7},
		want: false,
	},
	{
		name: "all distinct",
		s:    []int{1, 2, 3, 4},
		want: false,
	},
	{
		name: "duplicate at the start",
		s:    []int{1, 1, 2, 3, 4, 5, 6, 7},
		want: true,
	},

	{
		name: "duplicate at the end",
		s:    []int{1, 2, 3, 4, 5, 6, 7, 7},
		want: true,
	},
	{
		name: "distinct negative numbers",
		s:    []int{-1, -2, -3, -4, -5, -6, -7},
		want: false,
	},
	{
		name: "duplicate negative numbers",
		s:    []int{-1, -2, -3, -3, -4, -5, -6, -7},
		want: true,
	},
	{
		name: "distinct large values near the constraints",
		s: []int{1_000_000_000, 373158606, -760918337, -946289816, 592467581, -409379029,
			-474098744, -520658577, -700344588, 999_999_999, 581559892, -779893293,
			453201078, 590571865, 915941032, 171165722, -813300287,
			268073012, -93929780, -931747207, -936010953, -798790995,
			-530478523, -500364532, 85242216, 292825378, -943014439,
			205264595, -572999404, 537640408, 395616196, 506083910,
			170252923, -99088046, -526607376},
		want: false,
	},
	{
		name: "duplicate large values near the constraints",
		s: []int{-1_000_000_000, -779893293, -931747207, -798790995, -500364532, 506083910,
			-474098744, 292825378, -572999404, -530478523, -813300287,
			205264595, 268073012, -93929780, -700344588, -931747207,
			-760918337, 537640408, -936010953, 590571865, -93929780,
			915941032, -946289816, 592467581, 999_999_999, 395616196, -943014439,
			-520658577, -409379029, 581559892, -520658577, 171165722,
			85242216, 453201078, 373158606, -1_000_000_000},
		want: true,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, tc := range tests {

		sBefore := tc.s
		got := ContainsDuplicate(tc.s)

		if got != tc.want {
			t.Errorf("%s: expected %t, got %t", tc.name, tc.want, got)
		}

		if len(sBefore) != len(tc.s) {
			t.Errorf("Original slice is modified")
		} else {
			for i := 0; i < len(tc.s); i++ {
				if sBefore[i] != tc.s[i] {
					t.Errorf("Original slice is modified")
				}
			}
		}
	}

}

func TestContainsDuplicateSort(t *testing.T) {
	for _, tc := range tests {

		sBefore := tc.s
		got := ContainsDuplicateSort(tc.s)

		if got != tc.want {
			t.Errorf("%s: expected %t, got %t", tc.name, tc.want, got)
		}

		if len(sBefore) != len(tc.s) {
			t.Errorf("Original slice is modified")
		} else {
			for i := 0; i < len(tc.s); i++ {
				if sBefore[i] != tc.s[i] {
					t.Errorf("Original slice is modified")
				}
			}
		}

	}

}

func TestContainsDuplicateGeneric(t *testing.T) {
	for _, tc := range tests {

		sBefore := tc.s
		got := ContainsDuplicateGeneric(tc.s)

		if got != tc.want {
			t.Errorf("%s: expected %t, got %t", tc.name, tc.want, got)
		}

		if len(sBefore) != len(tc.s) {
			t.Errorf("Original slice is modified")
		} else {
			for i := 0; i < len(tc.s); i++ {
				if sBefore[i] != tc.s[i] {
					t.Errorf("Original slice is modified")
				}
			}
		}
	}

}
