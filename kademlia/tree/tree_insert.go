package tree

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
      t.correctInsertion(node)
      return true
    }
    return t.insert(node, root.left)
  }

  // When node.key > root.key
  if root.right == nil {
    root.right = node
    node.parent = root
    node.isLeftChild = false
    t.correctInsertion(node)
    return true
  }
  return t.insert(node, root.right)
}

func (t *Tree) correctInsertion(node *Node) {
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
    t.correctInsertion(grandParent)
    return
  }

  // Uncle is Black or Nil
  if node.isLeftChild {
    // Left-Left case, right-rotate
    if parent.isLeftChild {
      t.rightRotate(grandParent)
      t.correctInsertion(node)
      return
    }
    // Right-Left case, right-rotate
    t.rightLeftRotate(grandParent)
  } else {
    // Right-Right case, left-rotate
    if !parent.isLeftChild {
      t.leftRotate(grandParent)
      t.correctInsertion(node)
      return
    }
    // Left-Right case, left-rotate
    t.leftRightRotate(grandParent)
  }
}
