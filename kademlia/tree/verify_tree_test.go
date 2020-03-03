package tree

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tree Verify-Function", func() {
	var t *Tree

	BeforeEach(func() {
		t = &Tree{}
	})

	It("returns error if tree is nil", func() {
		Expect(verifyTree(nil)).To(HaveOccurred())
	})

	It("returns error if root is nil", func() {
		Expect(verifyTree(t)).To(HaveOccurred())
	})

	It("returns error if root is red", func() {
		t.Insert(NewNode(1, nil))
		t.FindNode(1, t.root).isBlack = false
		Expect(verifyTree(t)).To(HaveOccurred())
	})

	It("returns error on consecutive Red-Red Nodes", func() {
		t.Insert(NewNode(1, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(17, nil))
		t.Insert(NewNode(3, nil))
		t.Insert(NewNode(23, nil))
		t.Insert(NewNode(31, nil))
		t.Insert(NewNode(38, nil))
		t.Insert(NewNode(45, nil))

		Expect(verifyTree(t)).ToNot(HaveOccurred())

		// Node 45 is also Red,
		// so this creates Red-Red situation
		t.FindNode(38, t.root).isBlack = false

		Expect(verifyTree(t)).To(HaveOccurred())
	})

	It("returns error on unequal Black-nodes count in all paths", func() {
		t.Insert(NewNode(1, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(17, nil))
		t.Insert(NewNode(3, nil))
		t.Insert(NewNode(23, nil))
		t.Insert(NewNode(31, nil))
		t.Insert(NewNode(38, nil))
		t.Insert(NewNode(45, nil))

		Expect(verifyTree(t)).ToNot(HaveOccurred())

		// Node 17 is Black leaf-node,
		// so this invalidates the Black-count
		t.FindNode(23, t.root).right = nil
		Expect(verifyTree(t)).To(HaveOccurred())
	})

	It("returns error if a left node is higher than parent", func() {
		t.Insert(NewNode(1, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(17, nil))
		t.Insert(NewNode(3, nil))
		t.Insert(NewNode(23, nil))
		t.Insert(NewNode(31, nil))
		t.Insert(NewNode(38, nil))
		t.Insert(NewNode(45, nil))

		Expect(verifyTree(t)).ToNot(HaveOccurred())

		t.FindNode(23, t.root).left.key = 46
		Expect(verifyTree(t)).To(HaveOccurred())
	})

	It("returns error if a right node is lower than parent", func() {
		t.Insert(NewNode(1, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(17, nil))
		t.Insert(NewNode(3, nil))
		t.Insert(NewNode(23, nil))
		t.Insert(NewNode(31, nil))
		t.Insert(NewNode(38, nil))
		t.Insert(NewNode(45, nil))

		Expect(verifyTree(t)).ToNot(HaveOccurred())

		// The left-node is 17
		t.FindNode(23, t.root).key = 10
		Expect(verifyTree(t)).To(HaveOccurred())
	})

	It("returns error if duplicate keys are found", func() {
		t.Insert(NewNode(1, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(17, nil))
		t.Insert(NewNode(3, nil))
		t.Insert(NewNode(23, nil))
		t.Insert(NewNode(31, nil))
		t.Insert(NewNode(38, nil))
		t.Insert(NewNode(45, nil))

		Expect(verifyTree(t)).ToNot(HaveOccurred())

		t.FindNode(38, t.root).key = 15
		Expect(verifyTree(t)).To(HaveOccurred())
	})
})
