package tree

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Tree Verify-Function", func() {
  var t *Tree

  BeforeEach(func() {
    t = &Tree{}
    var _ = t
  })

  It("returns error if tree is nil", func() {
  	Expect(verifyTree(nil)).To(HaveOccurred())
  })

  It("returns error if root is nil", func() {
  	Expect(verifyTree(t)).To(HaveOccurred())
  })
})
