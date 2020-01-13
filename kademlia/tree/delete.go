package tree

import "github.com/pkg/errors"

func (t *Tree) Delete(node *Node) error {
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
    return t.Delete(predeccessor)
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
  // Parent is Red, and Sibling and Nephews are Nil/Black
  if !node.parent.isBlack && sibling.isBlack &&
    (sibling.left == nil || sibling.left.isBlack) &&
    (sibling.right == nil || sibling.right.isBlack) {
    // Exchange Parent and Sibling colors
    node.parent.isBlack = true
    sibling.isBlack = false
    return nil
  }

  // Parent is Red and Sibling is Black,
  // and one of the Nephews is Red
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
