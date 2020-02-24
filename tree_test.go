package RBtree

import "testing"

func TestTree_Insert(t1 *testing.T) {
	tree := NewTree()
	tree.Insert(50, 2)
	tree.Insert(45, 2)
	tree.Insert(55, 2)
	tree.Insert(40, 4)
	tree.Insert(47, 4)
	tree.Insert(35, 4)
	tree.Insert(42, 4)

	//tree.Insert(32, 4)
	tree.Delete(55)
	tree.Delete(45)

	tree.Print()
}
