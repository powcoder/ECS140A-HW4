https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.
type DisjointForest struct {
	parent map[int]int
}

func (forest *DisjointForest) FindSet(value int) int {
	if val, ok := forest.parent[value]; ok {
		if val == value {
			return value
		}
		forest.parent[value] = forest.FindSet(forest.parent[value])
		return forest.parent[value]
	}

	forest.parent[value] = value
	return value
}

func (forest *DisjointForest) UnionSet(a, b int) int {
	a = forest.FindSet(a)
	b = forest.FindSet(b)

	if a != b {
		forest.parent[b] = a
		return forest.FindSet(b)
	}

	return a
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
// This solution is adopted from my hw 1 solution
func NewDisjointSet() DisjointSet {
	var set DisjointForest
	set.parent = make(map[int]int)

	return &set
}
