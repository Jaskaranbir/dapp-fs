package tree

import (
  "fmt"
  "testing"

  "github.com/pkg/errors"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func TestRedBlackTree(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "RedBlackTree Suite")
}

// verifyTree checks if the root-tree is a valid Red-Black tree (development purposes).
// Following rules are tested:
// * Root-node is black
// * No two consecutive red-red nodes
// * Equal number of black nodes in all paths
// * Left-Node has lower key than parent
// * Right-node has higher key than parent
// * No duplicate keys
func verifyTree(t *Tree) error {
  if t.root == nil {
    err := errors.New("root is nil")
    return errors.Wrap(err, "failed validating tree")
  }

  type StackNode struct {
    node         *Node
    leftVisited  bool
    rightVisited bool
    blackCount   int
  }

  // Spicing things up with Stack;
  // because enough recursion!
  stack := []*StackNode{
    &StackNode{
      node:         t.root,
      leftVisited:  false,
      rightVisited: false,
      // Mankind has long known that root of RB-Tree is black
      blackCount: 1,
    },
  }

  // Keeps track of black-nodes in path(s)
  // (1 element = all black nodes in 1 path)
  blackCounter := make([]int, 0)

  for len(stack) != 0 {
    currStackNode := stack[len(stack)-1]

    node := currStackNode.node
    // Red-red nodes, invalid-condition
    if node.parent != nil && !node.parent.isBlack && !node.isBlack {
      err := fmt.Errorf("consecutive red-red nodes: %d with parent %d", node.key, node.parent.key)
      return errors.Wrap(err, "tree-validation failed")
    }

    // Detect leaf-nodes and store black-nodes counted in path
    if node != t.root {
      isRevisit := currStackNode.leftVisited || currStackNode.rightVisited

      if node.isBlack && !isRevisit {
        currStackNode.blackCount++
      }

      // Current node has atleast one leaf-node; we use Inorder-traversal:
      // If left-node is present, visit left-node(s) and then mark path is completed
      // If right-node is present, mark path as completed and visit right-node(s)
      if (!isRevisit && node.left == nil) || (isRevisit && node.right == nil) {
        blackCounter = append(blackCounter, currStackNode.blackCount)
      }
    }

    // Visit left if appliable
    if node.left != nil && !currStackNode.leftVisited {
      if node.left.key >= node.key {
        err := fmt.Errorf("left-node (%d) is >= right-node (%d)", node.left.key, node.key)
        return errors.Wrap(err, "tree-validation failed")
      }
      currStackNode.leftVisited = true

      stack = append(stack, &StackNode{
        node:         node.left,
        leftVisited:  false,
        rightVisited: false,
        blackCount:   currStackNode.blackCount,
      })
      continue
    }

    // Visit right if appliable
    if node.right != nil && !currStackNode.rightVisited {
      if node.right.key <= node.key {
        err := fmt.Errorf("right-node (%d) is <= left-node (%d)", node.right.key, node.key)
        return errors.Wrap(err, "tree-validation failed")
      }
      currStackNode.rightVisited = true

      stack = append(stack, &StackNode{
        node:         node.right,
        leftVisited:  false,
        rightVisited: false,
        blackCount:   currStackNode.blackCount,
      })
      continue
    }

    // Remove last element from slice if leaf
    // node found (or all children visited)
    stack = stack[:len(stack)-1]
  }

  for i := 1; i < len(blackCounter); i++ {
    // All paths should have equal black-nodes
    if blackCounter[i] != blackCounter[i-1] {
      err := fmt.Errorf("invalid black-count: %v", blackCounter)
      return errors.Wrap(err, "tree-validation failed")
    }
  }

  return nil
}
