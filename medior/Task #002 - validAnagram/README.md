# Task #002 - validAnagram

[Trello card](https://trello.com/c/YuypLHvP/2-task-002-validanagram)

## Overview
**Difficulty:** Easy
**Pattern:** Arrays & Hashing
**LeetCode reference:** [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)
**Repo folder:** `medior/Task #002 - validAnagram/task-002-valid-anagram/`

A small step up from Task #001: instead of detecting a repeat, you now compare the contents of two inputs. The focus is still clean, idiomatic Go and correct handling of edge cases.

## Problem
Given two strings `s` and `t`, return `true` if `t` is an **anagram** of `s`, and `false` otherwise.

An anagram is a word formed by rearranging the letters of another word, using every original letter **exactly once**.

## Function signature
```go
package validanagram

func IsAnagram(s string, t string) bool
```

## Examples
| `s` | `t` | Output | Why |
|---|---|---|---|
| `"anagram"` | `"nagaram"` | `true` | same letters, same counts |
| `"rat"` | `"car"` | `false` | different letters |
| `"aacc"` | `"ccac"` | `false` | same letters, different counts |
| `""` | `""` | `true` | two empty strings |

## Constraints
- `0 <= len(s), len(t) <= 5 * 10^4`
- `s` and `t` consist of **lowercase English letters** only (`a`-`z`).

## Requirements
1. Implement `IsAnagram` in O(n) time, where n is the length of the input.
2. Use O(1) extra space, taking advantage of the lowercase English letters constraint.
3. Also implement `IsAnagramUnicode(s string, t string) bool`, which must work correctly for **any Unicode text** (for example `"čaša"` and `"šača"` return `true`, and `"привет"` and `"тевирп"` return `true`).
4. Return `false` as early as possible when the inputs clearly cannot be anagrams.
5. Add a doc comment above each exported function.

## Bonus (optional)
- Write **table-driven tests** with `t.Run` subtests covering at least: both empty, one empty, different lengths, same letters with different counts, a valid anagram, identical strings, and Unicode inputs for `IsAnagramUnicode`. Run them with `-race`.
- In `README.md`, explain the time and space complexity of both functions and why the Unicode version can't use the same space trick (2 to 5 sentences each).
- Add benchmarks (`BenchmarkIsAnagram`, `BenchmarkIsAnagramUnicode`) on two strings of 50,000 characters and write the results in the README.
- Write a third version, `IsAnagramSort(s string, t string) bool`, that uses sorting, and note its complexity in the README.

## Expected project structure
```
task-002-valid-anagram/
├── go.mod
├── valid_anagram.go
├── valid_anagram_test.go   (bonus)
└── README.md               (bonus)
```

## Definition of Done
- [ ] `go build ./...` succeeds
- [ ] `go vet ./...` reports no issues
- [ ] `gofmt -l .` returns nothing
- [ ] Both functions return correct results (verified by my review tests, run with `-race`)
- [ ] `IsAnagramUnicode` handles multi-byte characters correctly

## Submission
Push your solution to the repo folder above, move this card to **Testing / QA / Review**, and send me a message that it is ready for review.
