package containsduplicate

import (
	"slices"
)

// ContainsDuplicate accept an integer slice nums,
// return true if any value appears at least twice in the slice,
// or false if every element is distinct.
func ContainsDuplicate(nums []int) bool {

	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {

		if _, hasValue := set[num]; hasValue {
			return true
		}

		set[num] = struct{}{}

	}

	return false
}

// ContainsDuplicateSort accept an integer slice nums,
// function then sort slice nums
// and return true if any value appears at least twice in the slice,
// or false if every element is distinct.
func ContainsDuplicateSort(nums []int) bool {

	numsClone := slices.Clone(nums)
	slices.Sort(numsClone)

	for i := 1; i < len(numsClone); i++ {
		if numsClone[i] == numsClone[i-1] {
			return true
		}
	}

	return false
}

// ContainsDuplicateGeneric accept an generic slice items,
// return true if any value appears at least twice in the slice,
// or false if every element is distinct.
func ContainsDuplicateGeneric[T comparable](items []T) bool {

	set := make(map[T]struct{}, len(items))

	for _, value := range items {
		if _, hasValue := set[value]; hasValue {
			return true
		}

		set[value] = struct{}{}
	}

	return false
}
