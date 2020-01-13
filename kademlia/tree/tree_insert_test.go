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
    Expect(err).Should(HaveOccurred())
  })

  It("sets first-node as root and sets node-properties", func() {
    node := NewNode(0, nil)
    node.isBlack = false
    node.tree = nil
    err := t.Insert(node)

    Expect(err).ShouldNot(HaveOccurred())
    Expect(t.root).To(Equal(node))
    Expect(node.isBlack).To(BeTrue())
    Expect(node.tree).To(Equal(t))
  })

  It("balances red-red", func() {
    err := t.Insert(NewNode(10, nil))
    Expect(err).ShouldNot(HaveOccurred())

    err = t.Insert(NewNode(-10, nil))
    Expect(err).ShouldNot(HaveOccurred())
    err = t.Insert(NewNode(-20, nil))
    Expect(err).ShouldNot(HaveOccurred())
    err = t.Insert(NewNode(6, nil))
    Expect(err).ShouldNot(HaveOccurred())
    err = t.Insert(NewNode(2, nil))
    Expect(err).ShouldNot(HaveOccurred())
    err = t.Insert(NewNode(4, nil))
    Expect(err).ShouldNot(HaveOccurred())

    err = t.Insert(NewNode(20, nil))
    Expect(err).ShouldNot(HaveOccurred())
    err = t.Insert(NewNode(15, nil))
    Expect(err).ShouldNot(HaveOccurred())
    err = t.Insert(NewNode(25, nil))
    Expect(err).ShouldNot(HaveOccurred())

    Expect(verifyTree(t)).To(Succeed())
  })

  It("returns error on duplicate insertion", func() {
    err := t.Insert(NewNode(10, nil))
    Expect(err).ShouldNot(HaveOccurred())

    err = t.Insert(NewNode(15, nil))
    Expect(err).ShouldNot(HaveOccurred())

    // Duplicate value
    err = t.Insert(NewNode(10, nil))
    Expect(err).Should(HaveOccurred())

    err = t.Insert(NewNode(20, nil))
    Expect(err).ShouldNot(HaveOccurred())

    Expect(verifyTree(t)).To(Succeed())
  })

  It("handles generic insertion", func() {
    t.Insert(NewNode(54, nil))
    t.Insert(NewNode(51, nil))
    t.Insert(NewNode(55, nil))
    t.Insert(NewNode(33, nil))
    t.Insert(NewNode(30, nil))
    t.Insert(NewNode(20, nil))
    t.Insert(NewNode(19, nil))
    t.Insert(NewNode(57, nil))
    t.Insert(NewNode(56, nil))
    t.Insert(NewNode(58, nil))
    t.Insert(NewNode(40, nil))
    t.Insert(NewNode(32, nil))
    t.Insert(NewNode(13, nil))
    t.Insert(NewNode(14, nil))
    t.Insert(NewNode(31, nil))
    t.Insert(NewNode(18, nil))
    t.Insert(NewNode(52, nil))
    t.Insert(NewNode(59, nil))
    t.Insert(NewNode(60, nil))
    t.Insert(NewNode(50, nil))
    t.Insert(NewNode(16, nil))
    t.Insert(NewNode(-3, nil))
    t.Insert(NewNode(100, nil))
    t.Insert(NewNode(53, nil))
    t.Insert(NewNode(92, nil))
    t.Insert(NewNode(70, nil))
    t.Insert(NewNode(12, nil))
    t.Insert(NewNode(11, nil))
    t.Insert(NewNode(15, nil))
    t.Insert(NewNode(192, nil))
    t.Insert(NewNode(200, nil))
    t.Insert(NewNode(204, nil))
    t.Insert(NewNode(1, nil))

    Expect(verifyTree(t)).To(Succeed())
  })
})
