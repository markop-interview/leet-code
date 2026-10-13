# Complexity

## Approach #1: ContainsDuplicate

`ContainsDuplicate` scans the slice once and records values already seen in a `map[int]struct{}`. Map lookup and insert are amortized O(1), so the expected running time is O(n). Extra memory is O(n) when every element is distinct; a duplicate found partway through returns immediately, so both time and space stay proportional to how far the scan got. Pick this approach for the usual case: `n` up to 10^5, a hash table fits in memory, and you want the early exit when a duplicate shows up before the end.

## Approach #2: ContainsDuplicateSort

`ContainsDuplicateSort` copies the caller's slice, sorts that copy with `sort.Ints`, then compares neighbors. The copy costs O(n) time and O(n) space, sorting costs O(n log n) time and O(log n) stack space, and the neighbor scan is O(n), so the function is O(n log n) time and O(n) space overall. The scan can stop at the first matching pair, but the sort always processes the full copy first. Pick this approach when you want a guaranteed O(n log n) bound, when a compact `[]int` copy is cheaper than a hash map's per-entry overhead, or when the input is already nearly sorted and that order makes the sort finish faster.
