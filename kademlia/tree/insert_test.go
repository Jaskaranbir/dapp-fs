package tree

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Tree Insert-Operation", func() {
  var t *Tree

  BeforeEach(func() {
    t = &Tree{}
  })

  It("returns error if node is nil", func() {
    err := t.Insert(nil)
    Expect(err).To(HaveOccurred())
  })

  It("sets first-node as root and sets node-properties", func() {
    node := NewNode(0, nil)
    node.isBlack = false
    node.tree = nil
    err := t.Insert(node)

    Expect(err).ToNot(HaveOccurred())
    Expect(t.root).To(Equal(node))
    Expect(node.isBlack).To(BeTrue())
    Expect(node.tree).To(Equal(t))
    Expect(t.size).To(Equal(1))
  })

  It("balances red-red", func() {
    err := t.Insert(NewNode(10, nil))
    Expect(err).ToNot(HaveOccurred())

    err = t.Insert(NewNode(-10, nil))
    Expect(err).ToNot(HaveOccurred())
    err = t.Insert(NewNode(-20, nil))
    Expect(err).ToNot(HaveOccurred())
    err = t.Insert(NewNode(6, nil))
    Expect(err).ToNot(HaveOccurred())
    err = t.Insert(NewNode(2, nil))
    Expect(err).ToNot(HaveOccurred())
    err = t.Insert(NewNode(4, nil))
    Expect(err).ToNot(HaveOccurred())

    err = t.Insert(NewNode(20, nil))
    Expect(err).ToNot(HaveOccurred())
    err = t.Insert(NewNode(15, nil))
    Expect(err).ToNot(HaveOccurred())
    err = t.Insert(NewNode(25, nil))
    Expect(err).ToNot(HaveOccurred())

    Expect(verifyTree(t)).To(Succeed())
    Expect(t.size).To(Equal(9))
  })

  It("returns error on duplicate insertion", func() {
    err := t.Insert(NewNode(10, nil))
    Expect(err).ToNot(HaveOccurred())

    err = t.Insert(NewNode(15, nil))
    Expect(err).ToNot(HaveOccurred())

    // Duplicate value
    err = t.Insert(NewNode(10, nil))
    Expect(err).To(HaveOccurred())

    err = t.Insert(NewNode(20, nil))
    Expect(err).ToNot(HaveOccurred())

    Expect(verifyTree(t)).To(Succeed())
    Expect(t.size).To(Equal(3))
  })

  It("handles generic insertion", func() {
    nodes := []*Node{
      NewNode(54, nil),
      NewNode(51, nil),
      NewNode(55, nil),
      NewNode(33, nil),
      NewNode(30, nil),
      NewNode(20, nil),
      NewNode(19, nil),
      NewNode(57, nil),
      NewNode(56, nil),
      NewNode(58, nil),
      NewNode(40, nil),
      NewNode(32, nil),
      NewNode(13, nil),
      NewNode(14, nil),
      NewNode(31, nil),
      NewNode(18, nil),
      NewNode(52, nil),
      NewNode(59, nil),
      NewNode(60, nil),
      NewNode(50, nil),
      NewNode(16, nil),
      NewNode(-3, nil),
      NewNode(100, nil),
      NewNode(53, nil),
      NewNode(92, nil),
      NewNode(70, nil),
      NewNode(12, nil),
      NewNode(11, nil),
      NewNode(15, nil),
      NewNode(192, nil),
      NewNode(200, nil),
      NewNode(204, nil),
      NewNode(1, nil),
    }
    for _, node := range nodes {
      err := t.Insert(node)
      Expect(err).ToNot(HaveOccurred())
    }

    Expect(verifyTree(t)).To(Succeed())
    Expect(t.size).To(Equal(len(nodes)))
  })
})
