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
		t.Insert(NewNode(10, nil))

		t.Insert(NewNode(-10, nil))
		t.Insert(NewNode(-20, nil))
		t.Insert(NewNode(6, nil))
		t.Insert(NewNode(2, nil))
		t.Insert(NewNode(4, nil))

		t.Insert(NewNode(20, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(25, nil))

		Expect(verifyTree(t)).To(Succeed())
	})

	It("returns false on duplicate insertion", func() {
		result := t.Insert(NewNode(10, nil))
		Expect(result).To(BeTrue())

		result = t.Insert(NewNode(15, nil))
		Expect(result).To(BeTrue())

		// Duplicate value
		result = t.Insert(NewNode(10, nil))
		Expect(result).To(BeFalse())

		result = t.Insert(NewNode(20, nil))
		Expect(result).To(BeTrue())

		Expect(verifyTree(t)).To(Succeed())
	})

	It("handles left-rotation", func() {
		t.Insert(NewNode(10, nil))

		t.Insert(NewNode(-10, nil))
		t.Insert(NewNode(7, nil))

		t.Insert(NewNode(20, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(13, nil))

		Expect(verifyTree(t)).To(Succeed())
	})

	It("handles right-rotation", func() {
		t.Insert(NewNode(10, nil))

		t.Insert(NewNode(-10, nil))
		t.Insert(NewNode(7, nil))

		t.Insert(NewNode(20, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(13, nil))

		Expect(verifyTree(t)).To(Succeed())
	})

	It("handles left-right-rotation", func() {
		t.Insert(NewNode(10, nil))

		t.Insert(NewNode(-10, nil))
		t.Insert(NewNode(7, nil))

		t.Insert(NewNode(20, nil))
		t.Insert(NewNode(15, nil))
		t.Insert(NewNode(17, nil))

		Expect(verifyTree(t)).To(Succeed())
	})

	It("handles right-left-rotation", func() {
		t.Insert(NewNode(10, nil))

		t.Insert(NewNode(20, nil))
		t.Insert(NewNode(17, nil))

		t.Insert(NewNode(-20, nil))
		t.Insert(NewNode(-15, nil))
		t.Insert(NewNode(-13, nil))

		Expect(verifyTree(t)).To(Succeed())
	})

	It("handles handles generic insertion", func() {
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
