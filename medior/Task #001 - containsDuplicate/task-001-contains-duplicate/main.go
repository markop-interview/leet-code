package main

import (
	"fmt"

	"github.com/markop-interview/leet-code/medior/task-001-contains-duplicate/containsduplicate"
)

func main() {

	testValues := []int{1, 2, 3, 1}
	test2Values := []int{1, 2, 3, 4}
	test3Values := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Printf("ContainsDuplicate test values %v: %t\n", testValues, containsduplicate.ContainsDuplicate(testValues))
	fmt.Printf("ContainsDuplicate test values %v: %t\n", test2Values, containsduplicate.ContainsDuplicate(test2Values))
	fmt.Printf("ContainsDuplicate test values %v: %t\n", test3Values, containsduplicate.ContainsDuplicate(test3Values))

	fmt.Printf("ContainsDuplicateSort test values %v: %t\n", testValues, containsduplicate.ContainsDuplicateSort(testValues))
	fmt.Printf("ContainsDuplicateSort test values %v: %t\n", test2Values, containsduplicate.ContainsDuplicateSort(test2Values))
	fmt.Printf("ContainsDuplicateSort test values %v: %t\n", test3Values, containsduplicate.ContainsDuplicateSort(test3Values))

	fmt.Printf("ContainsDuplicateGeneric test values %v: %t\n", testValues, containsduplicate.ContainsDuplicateGeneric(testValues))
	fmt.Printf("ContainsDuplicateGeneric test values %v: %t\n", test2Values, containsduplicate.ContainsDuplicateGeneric(test2Values))
	fmt.Printf("ContainsDuplicateGeneric test values %v: %t\n", test3Values, containsduplicate.ContainsDuplicateGeneric(test3Values))

}
