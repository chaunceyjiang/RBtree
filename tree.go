package RBtree

import (
	"fmt"
	"log"
)

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key int64, value interface{}) {
	node := NewNode(key, value)
	if t.root == nil {
		// The  root node is blcak
		node.color = BLACK
		t.root = node
	} else {
		t.insertNode(node)
	}
}

func (t *Tree) insertNode(node *Node) {
	pos, isLeft, exist := t.findNode(node.key)
	if exist {
		// update value
		pos.value = node.value
	} else if isLeft {
		node.parent = pos
		pos.left = node
	} else {
		node.parent = pos
		pos.right = node
	}
	t.insertFixes(node)
}

// insert fixes
func (t *Tree) insertFixes(curNode *Node) {
	if curNode.getParentNode() == nil {
		// 当前节点没有父节点，则表示该节点是 根节点，根节点必须是黑色
		curNode.setColor(BLACK)
		t.root = curNode
		return
	}
	// 1. 父亲节点是黑色, 则不修正，没有错误
	if curNode.getParentNode().getColor() == BLACK {
		return
	}
	// 2.1 父亲节点是红色并且叔叔节点也是红色
	if curNode.getUncleNode() != nil && curNode.getUncleNode().getColor() == RED {
		curNode.getParentNode().setColor(BLACK)
		curNode.getUncleNode().setColor(BLACK)
		curNode.getParentNode().getParentNode().setColor(RED)

		t.insertFixes(curNode.getGrandFather())
	} else {
		// 2.2 父亲节点是红色并且叔叔节点是黑色

		// 2.2.1
		curNode.getParentNode().setColor(BLACK)
		curNode.getGrandFather().setColor(RED)

		isLeft := curNode.getParentNode().left == curNode
		isParentLeft := curNode.getGrandFather().left == curNode.getParentNode()
		if isParentLeft && isLeft {
			gNode := curNode.getGrandFather()
			gNode.setColor(RED)
			curNode.getParentNode().setColor(BLACK)
			t.rotateRight(gNode)
		} else if isParentLeft && !isLeft {
			t.rotateLeft(curNode.getParentNode())
			gNode := curNode.getGrandFather()
			gNode.setColor(RED)
			curNode.getParentNode().setColor(BLACK)
			t.rotateRight(gNode)
		} else if !isParentLeft && !isLeft {
			gNode := curNode.getGrandFather()
			gNode.setColor(RED)
			curNode.getParentNode().setColor(BLACK)
			t.rotateLeft(gNode)
		} else {
			t.rotateRight(curNode.getParentNode())
			gNode := curNode.getGrandFather()
			gNode.setColor(RED)
			curNode.getParentNode().setColor(BLACK)
			t.rotateLeft(gNode)
		}
	}

}

func (t *Tree) Delete(key int64) {
	curNode, isLeft, exist := t.findNode(key)
	if !exist {
		return
	}
	// reNode 替换节点
	var reNode *Node
	if curNode.right != nil && curNode.left != nil {
		reNode = curNode.right.getSuffixNode()
		curNode.key = reNode.key
		curNode.value = reNode.value
	} else {
		reNode = curNode
	}
	t.realDelete(reNode, isLeft)
}
func (t *Tree) realDelete(curNode *Node, isLeft bool) {
	if curNode.left == nil && curNode.right == nil {
		if curNode.getColor() == RED {
			if isLeft {
				curNode.parent.left = nil
			} else {
				curNode.parent.right = nil
			}
			return
		}
		if curNode.parent == nil {
			t.root = nil
			return
		}
	}
	if (curNode.right != nil && curNode.left == nil) || (curNode.right == nil && curNode.left != nil) {
		switch isLeft {
		case true:
			curNode.parent.left = curNode.left
		case false:
			curNode.parent.right = curNode.right
		}
		return
	}
	t.deleteFixes(curNode)
	if isLeft {
		curNode.parent.left = curNode.left
	} else {
		curNode.parent.right = curNode.right
	}
}
func (t *Tree) deleteFixes(node *Node) {
	// 删除节点的兄弟节点是红色
	if node.getBrotherNode().getColor() == RED {
		node.getBrotherNode().setColor(BLACK)
		node.getParentNode().setColor(RED)
		isLeft := node.left == node.parent.left

		if isLeft {
			t.rotateLeft(node.getParentNode())
		} else {
			t.rotateRight(node.getParentNode())
		}
		t.deleteFixes(node)
	}
	if node.getBrotherNode().getColor() == BLACK {
		// nil Point error?
		switch node.getBrotherNode().right.getColor() {
		case RED:
			node.getBrotherNode().setColor(node.getParentNode().getColor())
			node.getParentNode().setColor(BLACK)
			node.getBrotherNode().right.setColor(BLACK)
			isLeft := node.left == node.parent.left
			if isLeft {
				t.rotateLeft(node.getParentNode())
			} else {
				t.rotateRight(node.getParentNode())
			}
		case BLACK:
			switch node.getBrotherNode().left.getColor() {
			case BLACK:
				isParentBlack := node.getParentNode().getColor() == BLACK
				node.getBrotherNode().setColor(RED)
				node.getParentNode().setColor(BLACK)
				if isParentBlack {
					t.deleteFixes(node.getParentNode())
				}
			case RED:
				node.getBrotherNode().setColor(RED)
				node.getBrotherNode().left.setColor(BLACK)
				isLeft := node.left == node.parent.left
				if isLeft {
					t.rotateRight(node.getBrotherNode())
				} else {
					t.rotateLeft(node.getBrotherNode())
				}
				t.deleteFixes(node)
			}
		}
	}
}

func (t *Tree) rotateLeft(node *Node) {
	t.rotate(RotateLeft, node)
}
func (t *Tree) rotateRight(node *Node) {
	t.rotate(RotateRight, node)
}
func (t *Tree) rotate(r RBrotate, node *Node) {
	root, err := node.nodeRotate(r)
	if err != nil {
		log.Fatal(err)
	}
	if root != nil {
		t.root = root
	}
}

func (t *Tree) findNode(key int64) (preNode *Node, isLeft bool, exist bool) {
	curNode := t.root
	for curNode != nil {
		if curNode.key == key {
			// no find
			return curNode, false, true
		}
		if curNode.key > key {
			preNode = curNode
			curNode = curNode.left
			isLeft = true
		} else {
			preNode = curNode
			curNode = curNode.right
			isLeft = false
		}
	}
	return
}

func (t *Tree) Print() {
	t.print(t.root)
}
func (t *Tree) print(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("key:%d value: %v ", node.key, node.value)
	if !node.color {
		fmt.Println("color: red")
	} else {
		fmt.Println("color: black")
	}
	t.print(node.left)
	t.print(node.right)
}

func NewTree() *Tree {
	return &Tree{root: nil}
}
