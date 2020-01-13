package tree

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Tree Find-Operations", func() {
  var t *Tree
  // Using map so its easier to select *Node
  var nodes map[int]*Node

  BeforeEach(func() {
    t = &Tree{}

    // It will help to use some RB-Tree-visualizer to see
    // node-arrangement for understanding test-cases better
    nodes = map[int]*Node{
      0:  NewNode(2, nil),
      1:  NewNode(34, nil),
      2:  NewNode(10, nil),
      3:  NewNode(1, nil),
      4:  NewNode(-8, nil),
      5:  NewNode(5, nil),
      6:  NewNode(50, nil),
      7:  NewNode(100, nil),
      8:  NewNode(80, nil),
      9:  NewNode(-2, nil),
      10: NewNode(12, nil),
      11: NewNode(56, nil),
      12: NewNode(97, nil),
      13: NewNode(48, nil),
      14: NewNode(23, nil),
      15: NewNode(13, nil),
      16: NewNode(75, nil),
      17: NewNode(67, nil),
      18: NewNode(53, nil),
      19: NewNode(78, nil),
      20: NewNode(83, nil),
      21: NewNode(42, nil),
      22: NewNode(47, nil),
      23: NewNode(101, nil),
      24: NewNode(82, nil),
    }

    // Push to array since order is important
    // (order affects how the tree is arranged)
    nodesArr := make([]*Node, len(nodes))
    for index, node := range nodes {
      nodesArr[index] = node
    }
    for _, node := range nodesArr {
      t.Insert(node)
    }
  })

  Describe("FindNode", func() {
    It("finds node under specified root", func() {
      result := t.FindNode(12, nodes[6])
      Expect(result.key).To(Equal(12))
    })

    It("finds node when root is tree-root", func() {
      result := t.FindNode(80, t.root)
      Expect(result.key).To(Equal(80))
    })

    It("doesn't search node outside root", func() {
      result := t.FindNode(50, nodes[17])
      Expect(result).To(BeNil())
    })

    It("returns nil when root-node is nil", func() {
      result := t.FindNode(80, nil)
      Expect(result).To(BeNil())
    })
  })

  Describe("Min", func() {
    It("finds min of specified root", func() {
      result := t.Min(nodes[8])
      Expect(result.key).To(Equal(53))
    })

    It("finds min when root has only right child", func() {
      result := t.Min(nodes[7])
      Expect(result.key).To(Equal(100))
    })

    It("finds min when node is root-node", func() {
      result := t.Min(t.root)
      Expect(result.key).To(Equal(-8))
    })

    It("finds min from root-node when node is nil", func() {
      result := t.Min(nil)
      Expect(result.key).To(Equal(-8))
    })

    It("returns root node when tree has root-node only", func() {
      t = &Tree{}
      root := NewNode(1, nil)
      t.Insert(root)

      result := t.Min(nil)
      Expect(result).To(Equal(root))
    })
  })

  Describe("Max", func() {
    It("finds max of specified root", func() {
      result := t.Max(nodes[8])
      Expect(result.key).To(Equal(101))
    })

    It("finds max when root has only left child", func() {
      result := t.Max(nodes[11])
      Expect(result.key).To(Equal(56))
    })

    It("finds max when node is root-node", func() {
      result := t.Max(t.root)
      Expect(result.key).To(Equal(101))
    })

    It("finds max from root-node when node is nil", func() {
      result := t.Max(nil)
      Expect(result.key).To(Equal(101))
    })

    It("returns root node when tree has root-node only", func() {
      t = &Tree{}
      root := NewNode(1, nil)
      t.Insert(root)

      result := t.Max(nil)
      Expect(result).To(Equal(root))
    })
  })

  Describe("InorderSuccessor", func() {
    It("finds inorder-successor of specified node", func() {
      result := t.InorderSuccessor(nodes[8])
      Expect(result.key).To(Equal(82))

      result = t.InorderSuccessor(nodes[17])
      Expect(result.key).To(Equal(75))

      result = t.InorderSuccessor(nodes[4])
      Expect(result.key).To(Equal(-2))
    })

    It("finds inorder-successor of root node", func() {
      result := t.InorderSuccessor(t.root)
      Expect(result.key).To(Equal(12))
    })

    It("returns nil if node has no inorder-successors", func() {
      result := t.InorderSuccessor(nodes[23])
      Expect(result).To(BeNil())
    })

    It("returns nil when node is nil", func() {
      result := t.InorderSuccessor(nil)
      Expect(result).To(BeNil())
    })

    It("returns nil when tree has root-node only", func() {
      t = &Tree{}
      root := NewNode(1, nil)
      t.Insert(root)

      result := t.InorderSuccessor(root)
      Expect(result).To(BeNil())
    })
  })

  Describe("InorderPredeccessor", func() {
    It("finds inorder-predeccessor of specified node", func() {
      result := t.InorderPredeccessor(nodes[8])
      Expect(result.key).To(Equal(78))

      result = t.InorderPredeccessor(nodes[17])
      Expect(result.key).To(Equal(56))

      result = t.InorderPredeccessor(nodes[9])
      Expect(result.key).To(Equal(-8))
    })

    It("finds inorder-predeccessor of root node", func() {
      result := t.InorderPredeccessor(t.root)
      Expect(result.key).To(Equal(5))
    })

    It("returns nil if node has no inorder-predeccessor", func() {
      result := t.InorderPredeccessor(nodes[4])
      Expect(result).To(BeNil())
    })

    It("returns nil when node is nil", func() {
      result := t.InorderPredeccessor(nil)
      Expect(result).To(BeNil())
    })

    It("returns nil when tree has root-node only", func() {
      t = &Tree{}
      root := NewNode(1, nil)
      t.Insert(root)

      result := t.InorderPredeccessor(root)
      Expect(result).To(BeNil())
    })
  })
})
