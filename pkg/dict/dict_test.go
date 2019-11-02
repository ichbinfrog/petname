package dict

import (
	"os"
	"testing"
)

func checkNodeValue(n *Node, k int, t *testing.T) {
	if n == nil {
		t.Errorf("Insert failed, node value is null")
	}

	if n.Key != k {
		t.Errorf("Insertion failed, expecting value %d", k)
	}
}

func TestDictInsert(t *testing.T) {
	tree := Tree{}

	// Test empty tree insert
	tree.Insert(10)
	checkNodeValue(tree.Root, 10, t)

	// Test non empty tree right insert
	tree.Insert(15)
	checkNodeValue(tree.Root.Right, 15, t)

	// Test non empty tree left insert
	tree.Insert(5)
	checkNodeValue(tree.Root.Left, 5, t)

	// Test higher depth l-left insert
	tree.Insert(2)
	checkNodeValue(tree.Root.Left.Left, 2, t)

	// Test higher depth l-right insert
	tree.Insert(7)
	checkNodeValue(tree.Root.Left.Right, 7, t)

	// Test higher depth r-left insert
	tree.Insert(12)
	checkNodeValue(tree.Root.Right.Left, 12, t)

	// Test higher depth r-right insert
	tree.Insert(17)
	checkNodeValue(tree.Root.Right.Right, 17, t)
}

func TestDictSearch(t *testing.T) {
	etree := &Tree{}

	// Test null entry
	if etree.Search([]int{1}) != false {
		t.Errorf("Search should fail when tree root is null")
	}

	// Test successful search
	tree := Tree{}
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	// Test depth 1 search
	if tree.Search([]int{10}) == false {
		t.Errorf("Search for [10] failed for root depth, should return true")
	}

	// Test depth 1 search failure
	if tree.Search([]int{0}) == true {
		t.Errorf("Search for [0] succeeded for root depth, should return false")
	}

	// Test higher len than depth search failure
	if tree.Search([]int{10, 15, 17, 20}) == true {
		t.Errorf("Search for [10,15,17,20] succeeded should return false due to higher length than depth")
	}

	// Test full search
	if tree.Search([]int{10, 15, 17}) == false {
		t.Errorf("Search for [10,15,17] failed should return true")
	}

	// Test full search failure
	if tree.Search([]int{10, 15, 0}) == true {
		t.Errorf("Search for [10,15,0] succeeded should return false")
	}
}

func TestDictPrint(t *testing.T) {
	tree := Tree{}
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	Print(os.Stdout, tree.Root, 2, '└')
}
