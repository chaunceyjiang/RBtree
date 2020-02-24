package RBtree

import "errors"

type RBcolor bool
type RBrotate bool

const (
	RED         RBcolor  = false
	BLACK       RBcolor  = true
	RotateLeft  RBrotate = true
	RotateRight RBrotate = false
)

type Node struct {
	key    int64
	value  interface{}
	left   *Node
	right  *Node
	parent *Node
	color  RBcolor
}

func (n *Node) getParentNode() *Node {
	return n.parent
}

func (n *Node) getBrotherNode() *Node {
	if n.getParentNode() == nil {
		return nil
	}

	if n.getParentNode().left == n {
		return n.getParentNode().right
	} else {
		return n.getParentNode().left
	}
}

func (n *Node) getGrandFather() *Node {
	return n.getParentNode().getParentNode()
}

func (n *Node) getColor() RBcolor {
	if n == nil{
		return BLACK
	}
	return n.color
}
func (n *Node) setColor(c RBcolor)  {
	n.color = c
}

func (n *Node) getUncleNode() *Node {
	if n.getGrandFather() == nil {
		return nil
	}
	if n.getGrandFather().left == n.parent {
		return n.getGrandFather().right
	} else {
		return n.getGrandFather().left
	}
}

// only rotate,not change color
func (n *Node) nodeRotate(rotate RBrotate) (*Node, error) {
	var root *Node
	if n == nil {
		return nil, nil
	}
	if rotate == RotateLeft && n.right == nil {
		return root, errors.New("左旋右节点不能为空")
	} else if rotate == RotateRight && n.left == nil {
		return root, errors.New("右旋左节点不能为空")
	}
	parent := n.parent
	var isleft bool

	if parent != nil {
		isleft = n.parent.left == n
	}
	// 右旋
	if rotate == RotateRight {
		grandson := n.left.right
		n.parent = n.left
		n.left.right = n
		n.left = grandson
	} else {  // 左旋
		grandson := n.right.left
		n.right.left = n
		n.parent = n.right
		n.right = grandson
	}

	if parent == nil {
		n.parent.parent = nil
		root = n.parent
	} else {
		if isleft {
			parent.left = n.parent
		} else {
			parent.right = n.parent
		}
		n.parent.parent = parent
	}
	return root, nil
}

// Replace Node
func (n *Node) getSuffixNode() *Node {
	if n.left == nil{
		return n
	}
	return n.left.getSuffixNode()
}

func NewNode(key int64, value interface{}) *Node {
	return &Node{
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
		color:  RED,
	}
}
