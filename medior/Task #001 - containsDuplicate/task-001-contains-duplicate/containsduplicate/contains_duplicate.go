package containsduplicate

import (
	"sort"
)

func ContainsDuplicate(nums []int) bool {

	// create set structure
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}

func ContainsDuplicateSort(nums []int) bool {

	sorted := make([]int, 0, len(nums))
	sorted = append(sorted, nums...)

	sort.Ints(sorted)
	//fmt.Printf("%v :: ", sorted)

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			return true
		}
	}

	return false
}
