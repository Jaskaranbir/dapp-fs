package tree

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "log"
)

var _ = Describe("Tree Delete-Operation", func() {
  var t *Tree
  // Using map so its easier to select specific *Node
  // (using map-value (node-key))
  var nodeKeys map[int]int

  BeforeEach(func() {
    t = &Tree{}

    // It will help to use some RB-Tree-visualizer to see
    // node-arrangement for understanding test-cases better
    nodeKeys = map[int]int{
      0:  2,
      1:  34,
      2:  10,
      3:  1,
      4:  8,
      5:  5,
      6:  50,
      7:  100,
      8:  80,
      9:  28,
      10: 12,
      11: 56,
      12: 97,
      13: 48,
      14: 23,
      15: 13,
      16: 75,
      17: 67,
      18: 53,
      19: 78,
      20: 83,
      21: 42,
      22: 47,
      23: 101,
      24: 82,
    }

    // Push to array since order is important
    // (order affects how the tree is arranged)
    nodesArr := make([]*Node, len(nodeKeys))
    for index, key := range nodeKeys {
      nodesArr[index] = NewNode(key, nil)
    }
    for _, node := range nodesArr {
      err := t.Insert(node)
      Expect(err).ToNot(HaveOccurred())
    }
    Expect(verifyTree(t)).To(Succeed())
    Expect(t.Size()).To(Equal(len(nodeKeys)))
  })

  It("returns error if node is nil", func() {
    err := t.Delete(nil)
    Expect(err).To(HaveOccurred())
    Expect(t.Size()).To(Equal(len(nodeKeys)))
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
    // Order to delete Nodes using keys
    keyTestOrder := []int{
      17,
      7,
      9,
      15,
      16,
      12,
      18,
      20,
      23,
      1,
      2,
      8,
      11,
      22,
      4,
      5,
      13,
      21,
      14,
      19,
      24,
      0,
      3,
      6,
      10,
    }

    for _, key := range keyTestOrder {
      nodeKey := nodeKeys[key]

      node := t.FindNode(nodeKey, t.root)
      err := t.Delete(node)
      if err != nil {
        log.Printf(
          "Errorred in deletion at Key: %d.\n" +
            "Attempted Key-order: %v",
          nodeKey,
          keyTestOrder,
        )
      }
      Expect(err).ToNot(HaveOccurred())
      Expect(t.FindNode(nodeKey, t.root)).To(BeNil())

      if t.size > 0 {
        Expect(verifyTree(t)).To(Succeed())
      } else {
        // Deleting Root-Node
        Expect(verifyTree(t)).To(HaveOccurred())
      }

      delete(nodeKeys, key)
      Expect(t.Size()).To(Equal(len(nodeKeys)))
    }
  })
})
