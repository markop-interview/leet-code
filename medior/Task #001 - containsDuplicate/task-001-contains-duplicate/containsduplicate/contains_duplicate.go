package containsduplicate

import (
	"slices"
)

// ContainsDuplicate reports whether nums contains any duplicate integer.
// It returns false when every element is distinct.
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

// ContainsDuplicateSort reports whether nums contains any duplicate integer.
// It sorts a copy of nums, leaving the input slice unchanged, and returns false
// when every element is distinct.
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

// ContainsDuplicateGeneric reports whether items contains any duplicate value.
// The element type must be comparable so values can be used as map keys.
// It returns false when every element is distinct.
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
