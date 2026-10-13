# leet-code

Private repo for Go interview practice. Each task is a LeetCode-style problem solved in idiomatic Go, then tested and reviewed before it counts as done.

Two tracks:

- [Medior](medior/README.md) — current track, about 150 tasks, tracked on [Go Interview Prep - Medior](https://trello.com/b/A0jOhpzi/go-interview-prep-medior).
- [Senior](senior/README.md) — later track, about 200–250 harder tasks, on a board named Go Interview Prep - Senior.

## Repository structure

```
.
├── README.md
├── medior/
│   ├── README.md
│   └── Task #001 - containsDuplicate/
│       ├── README.md
│       └── task-001-contains-duplicate/
│           ├── go.mod
│           ├── main.go
│           ├── README.md
│           └── containsduplicate/
│               ├── contains_duplicate.go
│               └── contains_duplicate_test.go
└── senior/
    └── README.md
```

## Folder and naming convention

```
<level>/Task #NNN - camelCaseName/README.md
<level>/Task #NNN - camelCaseName/task-NNN-kebab-name/
```

- The task `README.md` is the full problem statement: difficulty, pattern, signature, examples, constraints, requirements, and bonus work.
- `task-NNN-kebab-name/` is the Go module. `go.mod`, `main.go` (`package main`), and a `README.md` for complexity notes sit in that root. The solution lives in a package subdirectory (`<pkg>/<solution>.go` and `<pkg>/*_test.go`).
- `NNN` is a zero-padded number. Difficulty increases with the task number, starting at Task #001.

Example: `medior/Task #001 - containsDuplicate/task-001-contains-duplicate/`.

## Workflow and grading

Each task is a card on its Trello board. The card moves through:

1. **To-Do**
2. **In Progress**
3. **Testing / QA / Review**
4. **Done**

A reviewer runs the solution, including `go test -race`, and does a code review. The grade is written as:

`Total: X/10 (Requirements: Y/10; Bonus: Z/N)`

`X` and `Y` are out of 10. `Z/N` is the bonus score, and `N` depends on how many bonus items that task has. When the review finds issues, the card returns to **To-Do** with a bug checklist.

For roughly the first 30 tasks, table-driven tests and README complexity notes are optional bonus items.

## Checks

From the task module directory (`task-NNN-kebab-name/`):

```bash
go build ./...
go vet ./...
gofmt -l .
go test -race -cover ./...
```

`gofmt -l .` prints nothing when the package is formatted.

## Tracks

- [Medior](medior/README.md)
- [Senior](senior/README.md)

## Resources

- [NeetCode roadmap](https://neetcode.io/roadmap)
- [LeetCode](https://leetcode.com/)
- [Exercism Go track](https://exercism.org/tracks/go)
- [System Design Primer](https://github.com/donnemartin/system-design-primer)
- [ByteByteGo](https://bytebytego.com/)
- [Hello Interview](https://www.hellointerview.com/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://go.dev/tour/)
- [Pramp](https://www.pramp.com/)
