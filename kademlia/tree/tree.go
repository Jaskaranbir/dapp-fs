package tree

import (
	"github.com/pkg/errors"
	"log"
)

type Tree struct {
	root *Node
}

func (t *Tree) Insert(node *Node) bool {
	if node == nil {
		return false
	}
	if t.root == nil {
		node.isBlack = true
		t.root = node
		node.tree = t
		return true
	}

	result := t.insert(node, t.root)
	if result == true {
		node.tree = t
	}
	return result
}

func (t *Tree) insert(node *Node, root *Node) bool {
	if node.key == root.key {
		// Duplicate
		return false
	}

	if node.key < root.key {
		if root.left == nil {
			root.left = node
			node.parent = root
			node.isLeftChild = true
			t.correctTree(node)
			return true
		}
		return t.insert(node, root.left)
	}

	// When node.key > root.key
	if root.right == nil {
		root.right = node
		node.parent = root
		node.isLeftChild = false
		t.correctTree(node)
		return true
	}
	return t.insert(node, root.right)
}

func (t *Tree) correctTree(node *Node) {
	if node.parent == nil {
		node.isBlack = true
		return
	}
	if node.parent.isBlack {
		return
	}

	parent := node.parent
	grandParent := parent.parent
	uncle := grandParent.left
	if parent.isLeftChild {
		uncle = grandParent.right
	}

	if uncle != nil && !uncle.isBlack {
		parent.isBlack = true
		uncle.isBlack = true
		grandParent.isBlack = false
		t.correctTree(grandParent)
		return
	}

	// Uncle is Black or Nil
	if node.isLeftChild {
		// Left-Left case, right-rotate
		if parent.isLeftChild {
			t.rightRotate(grandParent)
			t.correctTree(node)
			return
		}
		// Right-Left case, right-rotate
		t.rightLeftRotate(grandParent)
	} else {
		// Right-Right case, left-rotate
		if !parent.isLeftChild {
			t.leftRotate(grandParent)
			t.correctTree(node)
			return
		}
		// Left-Right case, left-rotate
		t.leftRightRotate(grandParent)
	}
}

func (t *Tree) leftRotate(grandParent *Node) {
	child := grandParent.right
	grandParent.right = child.left
	if grandParent.right != nil {
		grandParent.right.parent = grandParent
		grandParent.right.isLeftChild = false
	}

	child.left = grandParent
	child.isLeftChild = grandParent.isLeftChild
	grandParent.isLeftChild = true
	child.parent = grandParent.parent
	grandParent.parent = child

	if child.parent == nil {
		t.root = child
	} else if child.isLeftChild {
		// In case of left-right rotate
		// Grandparent's parent is now "child", so parent of "child"
		child.parent.left = child
	} else {
		// In case of left rotate
		// Grandparent's parent is now "child", so parent of "child"
		child.parent.right = child
	}

	childColor := child.isBlack
	child.isBlack = grandParent.isBlack
	grandParent.isBlack = childColor
}

func (t *Tree) rightRotate(grandParent *Node) {
	child := grandParent.left
	grandParent.left = child.right
	if grandParent.left != nil {
		grandParent.left.parent = grandParent
		grandParent.left.isLeftChild = true
	}

	child.right = grandParent
	child.isLeftChild = grandParent.isLeftChild
	grandParent.isLeftChild = false
	child.parent = grandParent.parent
	grandParent.parent = child
	// child.isLeftChild = true

	if child.parent == nil {
		t.root = child
	} else if !child.isLeftChild {
		// In case of right-left rotate;
		// Grandparent's parent is now "child", so parent of "child"
		child.parent.right = child
	} else {
		// In case of right rotate;
		// Grandparent's parent is now "child", so parent of "child"
		child.parent.left = child
	}

	childColor := child.isBlack
	child.isBlack = grandParent.isBlack
	grandParent.isBlack = childColor
}

func (t *Tree) leftRightRotate(grandParent *Node) {
	t.leftRotate(grandParent.left)
	t.rightRotate(grandParent)
}

func (t *Tree) rightLeftRotate(grandParent *Node) {
	t.rightRotate(grandParent.right)
	t.leftRotate(grandParent)
}

func (t *Tree) Contains(key int) bool {
	return t.FindNode(key, t.root) != nil
}

func (t *Tree) String() {
	t.string(t.root)
}

func (t *Tree) string(node *Node) {
	if node == nil {
		return
	}

	if node.left != nil {
		t.string(node.left)
	}

	log.Println(node.key, node.isBlack)
	t.string(node.right)
}

func (t *Tree) FindNode(key int, root *Node) *Node {
	if root == nil {
		if t.root == nil {
			return nil
		}
		root = t.root
	}
	if key == root.key {
		return root
	}

	if key < root.key {
		// Left key is absent, so node doesn't exist
		if root.left == nil {
			return nil
		}
		return t.FindNode(key, root.left)
	}

	// Right key is absent, so node doesn't exist
	if root.right == nil {
		return nil
	}
	return t.FindNode(key, root.right)
}

func (t *Tree) Min(root *Node) *Node {
	if root == nil {
		root = t.root
	}

	if root.left == nil {
		return root
	}
	return t.Min(root.left)
}

func (t *Tree) Max(root *Node) *Node {
	if root == nil {
		root = t.root
	}

	if root.right == nil {
		return root
	}
	return t.Max(root.right)
}

func (t *Tree) InorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return t.Min(node.right)
	}

	// Last left-ancestor
	lastLeftNode := node.parent
	for lastLeftNode.isLeftChild && lastLeftNode.parent != nil {
		lastLeftNode = lastLeftNode.parent
	}

	// The passed Node is largest Node in Tree
	if lastLeftNode.key < node.key {
		return nil
	}
	return lastLeftNode
}

func (t *Tree) InorderPredeccessor(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return t.Max(node.left)
	}

	// Last right-ancestor
	lastRightNode := node.parent
	for lastRightNode.isLeftChild && lastRightNode.parent != nil {
		lastRightNode = lastRightNode.parent
	}

	if node.isLeftChild {
		return lastRightNode.parent
	}
	return lastRightNode
}

func (t *Tree) Remove(node *Node) error {
	if node == nil {
		err := errors.New("nil node")
		return errors.Wrap(err, "error removing node")
	}
	if node.tree != t {
		err := errors.New("node doesn't belong to this tree")
		return errors.Wrap(err, "error removing node")
	}

  // No Children
	if node.left == nil && node.right == nil {
		// All nodes deleted from Tree :(
		if node == t.root {
			t.root = nil
			return nil
		}

		if node.isLeftChild {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		if node.isBlack {
			// Extra Black-node case
			return t.correctDeletion(node)
		}

		return nil
	}

	// Two Children
	if node.left != nil && node.right != nil {
		predeccessor := t.InorderPredeccessor(node)
		node.key = predeccessor.key
		node.value = predeccessor.value
		return t.Remove(predeccessor)
	}

	// Only Left-Child
	if node.left != nil {
		node.key = node.left.key
		node.value = node.left.value
		// Black-node with Black-child
		isExtraBlack := node.isBlack && node.left.isBlack
		node.left = nil
		if isExtraBlack {
			return t.correctDeletion(node)
		}
		return nil
	}

	// Only Right-Child
	if node.right != nil {
		successor := t.InorderSuccessor(node)
		node.key = successor.key
		node.value = successor.value

		if successor.left == nil && successor.right == nil {
			if !successor.isBlack {
				if successor.isLeftChild {
					successor.parent.left = nil
				} else {
					successor.parent.right = nil
				}
				return nil
			}

			if successor.isBlack {
				if !node.isBlack {
					node.isBlack = true
					return nil
				}
				return t.correctDeletion(node)
			}
		}

		if successor.right != nil {
			successor.parent.left = successor.right
			if !successor.right.isBlack {
				successor.right.isBlack = successor.isBlack
				return nil
			}
		}
		// Successor and its child are black,
		// or Successor is a Black leaf-node
		return t.correctDeletion(node)
	}

	return nil
}

func (t *Tree) correctDeletion(node *Node) error {
	sibling := node.Sibling()
	// Parent is Red, and Sibling and Nephews are Null/Black
	if !node.parent.isBlack && sibling.isBlack &&
		(sibling.left == nil || sibling.left.isBlack) &&
		(sibling.right == nil || sibling.right.isBlack) {
		// Exchange Parent and Sibling colors
		node.parent.isBlack = true
		sibling.isBlack = false
		return nil
	}

	// Parent is Red and Sibling is Black,
	// and one of the Newphews is Red
	if !node.parent.isBlack && sibling.isBlack {
		// Node is left, Sibling has Red-Right-child
		if node.isLeftChild && sibling.right != nil && !sibling.right.isBlack {
			t.leftRotate(node.parent)
			sibling.right.isBlack = true
			return nil
		}

		// Node is Right, Sibling has Red-Left-child
		if !node.isLeftChild && sibling.left != nil && !sibling.left.isBlack {
			t.rightRotate(node.parent)
			sibling.left.isBlack = true
			return nil
		}
	}

	if node.parent.isBlack && sibling.isBlack {
		// Sibling has No-child or two Black-children
		if (sibling.right == nil && sibling.left == nil) ||
			 (sibling.left != nil && sibling.right != nil &&
				sibling.left.isBlack && sibling.right.isBlack) {
			// Mark Sibling Red and push Double-Black to Parent
			sibling.isBlack = false
			parent := sibling.parent

			// Root is double-black, make it single-Black and be done
			if parent == t.root {
				return nil
			}
			return t.correctDeletion(sibling.parent)
		}

		// Left-Sibling
		if sibling.isLeftChild {
			// Sibling has Left-Red-child
			if sibling.left != nil && !sibling.left.isBlack {
				t.rightRotate(node.parent)
				sibling.left.isBlack = true
				return nil
			}

			// Sibling has Right-Red-child
			if sibling.right != nil && !sibling.right.isBlack {
				t.leftRightRotate(node.parent)
				sibling.isBlack = true
				return nil
			}
		}

		// Right-Sibling
		if !sibling.isLeftChild {
			// Sibling has Right-Red-child
			if sibling.right != nil && !sibling.right.isBlack {
				t.leftRotate(node.parent)
				sibling.right.isBlack = true
				return nil
			}

			// Sibling has Left-Red-child
			if sibling.left != nil && !sibling.left.isBlack {
				t.rightLeftRotate(node.parent)
				sibling.isBlack = true
				return nil
			}
		}
	}

	if node.parent.isBlack && !sibling.isBlack {
		if sibling.isLeftChild {
			t.rightRotate(sibling.parent)
		} else {
			t.leftRotate(sibling.parent)
		}
		return t.correctDeletion(node)
	}

	return nil
}
