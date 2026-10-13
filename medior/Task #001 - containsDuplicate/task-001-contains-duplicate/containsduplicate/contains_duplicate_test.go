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

func TestContainDuplicate(t *testing.T) {

	tests := []struct {
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
			s:    []int{1, 2, 3, 4, 5, 6, 7},
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
			name: "negative numbers",
			s:    []int{1, -2, -3, 4, 3, 5, -6, 7, -7},
			want: false,
		},
		{
			name: "duplicate negative numbers",
			s:    []int{1, -2, -3, 4, 3, 5, -6, 7, -7, -7},
			want: true,
		},
		{
			name: "large values with duplicates near the constraints",
			s: []int{1000000000, -1000000000, 42, 42, 0, -999999999, 999999999, 17, -17, 42,
				500000000, -500000000, 1, 1, -1, 987654320, -987654320, 999999998, -999999998, 100,
				-100, 0, 42, -42, 765432109, -765432109, 123456789, -123456789, 7, 7,
				-7, 1000000000, -1000000000},
			want: true,
		},
		{
			name: "large values without duplicates near the constraints",
			s: []int{1000000000, -1000000000, 999999999, -999999999, 999999998, -999999998, 876543210, -876543210, 765432109, -765432109,
				654321098, -654321098, 543210987, -543210987, 432109876, -432109876, 321098765, -321098765, 987654320, -987654320,
				123456789, -123456789, 987654321, -987654321, 100000001, -100000001, 500000000, -500000000, 314159265, -314159265,
				271828182, -271828182, 0},
			want: false,
		},
	}

	for _, tt := range tests {

		got := ContainsDuplicate(tt.s)

		if got != tt.want {
			t.Fatalf(
				"%s: expected %t, got %t",
				tt.name,
				tt.want,
				got,
			)
		}

	}

}

func TestContainDuplicateSort(t *testing.T) {
	tests := []struct {
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
			s:    []int{1, 2, 3, 4, 5, 6, 7},
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
			name: "negative numbers",
			s:    []int{1, -2, -3, 4, 3, 5, -6, 7, -7},
			want: false,
		},
		{
			name: "duplicate negative numbers",
			s:    []int{1, -2, -3, 4, 3, 5, -6, 7, -7, -7},
			want: true,
		},
		{
			name: "large values with duplicates near the constraints",
			s: []int{1000000000, -1000000000, 42, 42, 0, -999999999, 999999999, 17, -17, 42,
				500000000, -500000000, 1, 1, -1, 987654320, -987654320, 999999998, -999999998, 100,
				-100, 0, 42, -42, 765432109, -765432109, 123456789, -123456789, 7, 7,
				-7, 1000000000, -1000000000},
			want: true,
		},
		{
			name: "large values without duplicates near the constraints",
			s: []int{1000000000, -1000000000, 999999999, -999999999, 999999998, -999999998, 876543210, -876543210, 765432109, -765432109,
				654321098, -654321098, 543210987, -543210987, 432109876, -432109876, 321098765, -321098765, 987654320, -987654320,
				123456789, -123456789, 987654321, -987654321, 100000001, -100000001, 500000000, -500000000, 314159265, -314159265,
				271828182, -271828182, 0},
			want: false,
		},
	}

	for _, tt := range tests {

		got := ContainsDuplicateSort(tt.s)

		if got != tt.want {
			t.Fatalf(
				"%s: expected %t, got %t",
				tt.name,
				tt.want,
				got,
			)
		}

	}

}
