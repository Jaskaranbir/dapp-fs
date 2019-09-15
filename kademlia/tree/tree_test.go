package tree

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
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

var _ = Describe("RedBlackTree", func() {

	var t *Tree

	BeforeEach(func() {
		t = &Tree{}
	})

	Describe("insert", func() {
		insertAndVerify := func(key int) {
			result := t.Insert(NewNode(key, nil))
			Expect(result).To(BeTrue())
			Expect(verifyTree(t)).To(Succeed())
		}

		It("returns false if node is nil", func() {
			result := t.Insert(nil)
			Expect(result).To(BeFalse())
		})

		It("sets first-node as root and sets node-properties", func() {
			node := NewNode(0, nil)
			node.isBlack = false
			node.tree = nil
			result := t.Insert(node)

			Expect(result).To(BeTrue())
			Expect(t.root).To(Equal(node))
			Expect(node.isBlack).To(BeTrue())
			Expect(node.tree).To(Equal(t))
		})

		It("balances red-red", func() {
			insertAndVerify(10)

			insertAndVerify(-10)
			insertAndVerify(-20)
			insertAndVerify(6)
			insertAndVerify(2)
			insertAndVerify(4)

			insertAndVerify(20)
			insertAndVerify(15)
			insertAndVerify(25)
		})

		It("handles left-rotation", func() {
			insertAndVerify(10)

			insertAndVerify(-10)
			insertAndVerify(7)

			insertAndVerify(20)
			insertAndVerify(15)
			insertAndVerify(13)
		})

		It("handles right-rotation", func() {
			insertAndVerify(10)

			insertAndVerify(-10)
			insertAndVerify(7)

			insertAndVerify(20)
			insertAndVerify(15)
			insertAndVerify(13)
		})

		It("handles left-right-rotation", func() {
			insertAndVerify(10)

			insertAndVerify(-10)
			insertAndVerify(7)

			insertAndVerify(20)
			insertAndVerify(15)
			insertAndVerify(17)
		})

		It("handles right-left-rotation", func() {
			insertAndVerify(10)

			insertAndVerify(20)
			insertAndVerify(17)

			insertAndVerify(-20)
			insertAndVerify(-15)
			insertAndVerify(-13)
		})
	})
})
