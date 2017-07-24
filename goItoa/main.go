package main

import "fmt"

type conflictType int

const (
	conflictDependentChild conflictType = (1 << iota)
	conflictRunningContainer
	conflictActiveReference
	conflictStoppedContainer
	conflictHard = conflictDependentChild | conflictRunningContainer
	conflictSoft = conflictActiveReference | conflictStoppedContainer
)

func main() {
	fmt.Printf("%d %d %d %d %d %d\n",
		conflictDependentChild, conflictRunningContainer,
		conflictActiveReference, conflictStoppedContainer,
		conflictHard, conflictSoft)

	fmt.Printf("%b %b %b %b %b %b\n",
		conflictDependentChild, conflictRunningContainer,
		conflictActiveReference, conflictStoppedContainer,
		conflictHard, conflictSoft)
}
