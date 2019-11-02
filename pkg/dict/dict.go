// Package dict encapsulates the structure which is in charge
// of populating the petname array
package dict

import (
	"fmt"
	"io"
)

// Node represents a node in the Used binary
// tree to accelerate search speed
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Tree is a representation of a binary search tree
type Tree struct {
	Root  *Node
	Depth int
}

// Insert inserts a key into a binary tree
func (t *Tree) Insert(k interface{}) {
	switch k.(type) {
	case int:
		if t.Root == nil {
			t.Root = &Node{
				Key:   k.(int),
				Left:  nil,
				Right: nil,
			}
			t.Depth = 1
		} else {
			t.Root.insert(k.(int))
			t.Depth++
		}
	case []int:
		for _, v := range k.([]int) {
			t.Insert(v)
		}
	}
}

// insert inserts a key from the given root node
func (n *Node) insert(k int) {
	if n == nil {
		return
	} else if k <= n.Key {
		if n.Left == nil {
			n.Left = &Node{
				Key:   k,
				Left:  nil,
				Right: nil,
			}
		} else {
			n.Left.insert(k)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Key:   k,
				Left:  nil,
				Right: nil,
			}
		} else {
			n.Right.insert(k)
		}
	}
}

// Search finds a consecutive series of int in the binary tree
func (t *Tree) Search(k []int) bool {
	if t == nil || t.Root == nil || len(k) > t.Depth {
		return false
	}

	tmp := k
	node := t.Root
	for {
		if node == nil || (node.Key != tmp[0]) {
			return false
		}

		if len(tmp) <= 1 {
			return true
		}
		tmp = tmp[1:len(tmp)]
		if tmp[0] > node.Key {
			node = node.Right
		} else {
			node = node.Left
		}
	}
}

// Print draws the current tree (only useful for debugging)
func Print(w io.Writer, n *Node, ns int, ch rune) {
	if n == nil {
		return
	}

	fmt.Fprintf(w, "%*c── %v\n", ns+1, ch, n.Key)
	if n.Right != nil {
		Print(w, n.Left, ns+4, '├')
	} else {
		Print(w, n.Left, ns+4, '└')
	}
	Print(w, n.Right, ns+4, '└')
}
