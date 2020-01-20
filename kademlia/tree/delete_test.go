package tree

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Tree Delete-Operation", func() {
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
    Expect(verifyTree(t)).To(Succeed())
    Expect(t.size).To(Equal(len(nodes)))
  })

  It("returns error if node is nil", func() {
    err := t.Delete(nil)
    Expect(err).To(HaveOccurred())
    Expect(t.size).To(Equal(len(nodes)))
  })

  It("returns error if node belongs to different tree", func() {
		t2 := &Tree{}
		node := NewNode(12, nil)
		err := t2.Insert(node)
    Expect(err).ToNot(HaveOccurred())

    err = t.Delete(node)
    Expect(err).To(HaveOccurred())
    Expect(t2.size).To(Equal(1))
  })

  It("deletes root-node when tree has only root-node", func() {
		t2 := &Tree{}
		node := NewNode(12, nil)
		err := t2.Insert(node)
    Expect(err).ToNot(HaveOccurred())

    err = t2.Delete(node)
    Expect(err).ToNot(HaveOccurred())
    Expect(t2.size).To(BeZero())
  })

  It("returns error when node does not exist", func() {
    err := t.Delete(NewNode(200, nil))
    Expect(err).To(HaveOccurred())
	})

	It("handles generic deletion", func() {
		for k, v := range nodes {
			t.Delete(v)
			delete(nodes, k)
			Expect(verifyTree(t)).To(Succeed())
			Expect(t.size).To(Equal(len(nodes)))
		}
	})
})
