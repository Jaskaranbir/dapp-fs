package tree

import (
	"github.com/Jaskaranbir/dapp-fs/kademlia"
)

type Node struct {
	key   int
	value *kademlia.Node

	tree				*Tree
	isBlack     bool
	isLeftChild bool
	parent      *Node
	left        *Node
	right       *Node
}

func NewNode(key int, value *kademlia.Node) *Node {
	return &Node{
		key:   key,
		value: value,

		// Red by default
		isBlack:     false,
		isLeftChild: false,
		left:        nil,
		right:       nil,
	}
}

func (n *Node) Key() int {
	return n.key
}

func (n *Node) Value() *kademlia.Node {
	return n.value
}

// Tree this Node belongs to, if any.
func (n *Node) Tree() *Tree {
	return n.tree
}

func (n *Node) IsBlack() bool {
	return n.isBlack
}

func (n *Node) IsLeftChild() bool {
	return n.isLeftChild
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) Sibling() *Node {
	if n.parent == nil {
		return nil
	}
	if n.isLeftChild {
		return n.parent.right
	}
	return n.parent.left
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}
