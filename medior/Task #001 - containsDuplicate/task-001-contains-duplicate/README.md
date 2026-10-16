# Complexity

## Approach #1: ContainsDuplicate

`ContainsDuplicate` scans the slice once and records values already seen in a `map[int]struct{}`; expected time is $O(n)$ because map lookups and insertions are amortized $O(1)$. It uses $O(n)$ extra space in the all-distinct case, though it can stop early and use space proportional to the scanned prefix when it finds a duplicate. Choose it for a straightforward, fast general-purpose solution when the extra hash-map memory is acceptable and early detection is useful.

## Approach #2: ContainsDuplicateSort

`ContainsDuplicateSort` copies the input, sorts the copy with `slices.Sort`, and checks adjacent values, for $O(n \log n)$ time overall and $O(n)$ extra space for the copy (the sort itself uses $O(\log n)$ stack space). It preserves the caller's slice, but must finish sorting before the scan can find a duplicate. Choose it when predictable sorting bounds or avoiding hash-map overhead matter more than the extra copy and slower asymptotic time.

## Approach #3: ContainsDuplicateGeneric

`ContainsDuplicateGeneric` uses the same seen-value map strategy as `ContainsDuplicate`, but accepts any comparable element type; expected time is $O(n)$ and extra space is $O(n)$. It can return early when it encounters a repeated value, making actual work and map space proportional to the prefix scanned. Choose it when duplicate checking is needed for types other than `int` (such as strings or comparable structs); use `ContainsDuplicate` for integer slices when a type-specific API is preferred.
