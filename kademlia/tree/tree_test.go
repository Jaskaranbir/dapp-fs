package tree

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedBlackTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedBlackTree Suite")
}

var _ = Describe("Tree Generic-Operations", func() {
	var t *Tree

	BeforeEach(func() {
		t = &Tree{}
	})

	Describe("Tree Rotations", func() {
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
	})
})
