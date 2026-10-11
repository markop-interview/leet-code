## Overview
**Difficulty:** Easy (warm-up)
**Pattern:** Arrays & Hashing
**LeetCode reference:** [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

This is the warm-up task. The algorithm is simple on purpose, so the focus is on clean, idiomatic Go, just like in a real interview.

## Problem
Given an integer slice `nums`, return `true` if any value appears **at least twice** in the slice, and `false` if every element is distinct.

## Function signature
```go
package containsduplicate

func ContainsDuplicate(nums []int) bool
```

## Examples
| Input | Output | Why |
|---|---|---|
| `[1, 2, 3, 1]` | `true` | `1` appears twice |
| `[1, 2, 3, 4]` | `false` | all elements are distinct |
| `[1, 1, 1, 3, 3, 4, 3, 2, 4, 2]` | `true` | several duplicates |

## Constraints
- `0 <= len(nums) <= 10^5`
- `-10^9 <= nums[i] <= 10^9`
- The input slice must **not** be modified by your function.

## Requirements
1. Implement an **optimal** solution in O(n) time.
2. Also implement a second approach, `ContainsDuplicateSort(nums []int) bool`, that uses sorting (O(n log n) time). It must not mutate the caller's slice, so work on a copy.
3. Use idiomatic Go: a `map[int]struct{}` as a set is preferred over `map[int]bool`, and pre-size the map where it makes sense.
4. Return early as soon as a duplicate is found.
5. Add a doc comment above each exported function.

## Bonus (optional)
- Write **table-driven tests** with `t.Run` subtests covering at least: empty slice, single element, all distinct, duplicate at the start, duplicate at the end, negative numbers, and large values near the constraints. Both implementations should pass the same test table.
- In `README.md`, explain the time and space complexity of both approaches and when you would pick one over the other (2 to 5 sentences each).
- Add benchmarks (`BenchmarkContainsDuplicate`, `BenchmarkContainsDuplicateSort`) on a slice of 100,000 distinct elements and write the results in the README.
- Write a generic version: `func ContainsDuplicateGeneric[T comparable](items []T) bool`.

## Expected project structure
```
task-001-contains-duplicate/
├── go.mod
├── contains_duplicate.go
├── contains_duplicate_test.go   (bonus)
└── README.md                    (bonus)
```

## Definition of Done
- [ ] `go build ./...` succeeds
- [ ] `go vet ./...` reports no issues
- [ ] `gofmt -l .` returns nothing
- [ ] Both functions return correct results (verified by my review tests, run with `-race`)
- [ ] The input slice is never mutated

## Submission
Move this card to **Testing / QA / Review**, upload the project as a `.rar` in chat, and send a message that it is ready for review.
