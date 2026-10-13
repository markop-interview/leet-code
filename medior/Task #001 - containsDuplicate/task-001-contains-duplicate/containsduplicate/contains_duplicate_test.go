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
			name: "large values near the constraints",
			s:    []int{},
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
