package containsduplicate

import (
	"slices"
)

/*
Given an integer slice nums, return true if any value appears at least twice in the slice, and false if every element is distinct.
*/
func ContainsDuplicate(nums []int) bool {

	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {

		// if val, ok := scores["Bob"]; ok {
		if _, hasValue := set[num]; hasValue {
			return true
		}

		set[num] = struct{}{}

	}

	return false
}

/*
Given an integer slice nums, function then sort slice nums and return true if any value appears at least twice in the slice, and false if every element is distinct.
*/
func ContainsDuplicateSort(nums []int) bool {

	copy := slices.Clone(nums)
	slices.Sort(copy)

	for i := 1; i < len(nums); i++ {
		if copy[i] == copy[i-1] {
			return true
		}
	}

	return false
}
