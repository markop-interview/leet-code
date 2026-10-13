package main

import (
	"fmt"

	"github.com/markop-interview/leet-code/containsduplicate"
)

func main() {

	/*
		test2 := []int{1, 2, 3, 4}
		test3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	*/

	slice1 := make([]int, 0, 4)
	slice2 := make([]int, 0, 4)
	slice3 := make([]int, 0, 10)
	slice4 := make([]int, 0, 10)

	slice1 = []int{1, 2, 3, 1}
	slice2 = []int{1, 2, 3, 4}
	slice3 = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	slice4 = []int{1, 3, 8, 9, 7, 5, 4, 6, 10, 10}

	fmt.Println("[1, 2, 3, 1] ->", containsduplicate.ContainsDuplicate(slice1))
	fmt.Println("[1, 2, 3, 4] ->", containsduplicate.ContainsDuplicate(slice2))
	fmt.Println("[1, 1, 1, 3, 3, 4, 3, 2, 4, 2] ->", containsduplicate.ContainsDuplicate(slice3))
	fmt.Println("[1, 3, 8, 9, 7, 5, 4, 6, 10, 10] ->", containsduplicate.ContainsDuplicate(slice4))

	fmt.Println("[1, 2, 3, 1] ->", containsduplicate.ContainsDuplicateSort(slice1))
	fmt.Println("[1, 2, 3, 4] ->", containsduplicate.ContainsDuplicateSort(slice2))
	fmt.Println("[1, 1, 1, 3, 3, 4, 3, 2, 4, 2] ->", containsduplicate.ContainsDuplicateSort(slice3))
	fmt.Println("[1, 3, 8, 9, 7, 5, 4, 6, 10, 10] ->", containsduplicate.ContainsDuplicateSort(slice4))
}
