package tree

import (
  "fmt"
)

type Tree struct {
  root *Node
}

func (t *Tree) String(node *Node) {
  t.string(node)
}

func (t *Tree) string(node *Node) {
  if node == nil {
    return
  }

  if node.left != nil {
    t.string(node.left)
  }

  fmt.Println(node.key, node.isBlack)
  t.string(node.right)
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
