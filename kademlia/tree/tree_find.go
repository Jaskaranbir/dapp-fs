package tree

func (t *Tree) FindNode(key int, root *Node) *Node {
  if root == nil {
    return nil
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
  if node.parent == nil {
    return nil
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
  if node.parent == nil {
    return nil
  }

  // Find Last right-ancestor
  lastRightNode := node.parent
  for !lastRightNode.isLeftChild && lastRightNode.parent != nil {
    lastRightNode = lastRightNode.parent
  }

  if lastRightNode.key > node.key {
    return nil
  }
  return lastRightNode
}
