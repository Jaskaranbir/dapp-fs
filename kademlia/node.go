package kademlia

import (
	"time"
)

type Node struct {
	id   ID
	addr string

	lastSeen   int64
	staleCount int
}

func NewNode() (*Node, error) {
	node := &Node{
		staleCount: 0,
	}
	node.UpdateSeen()
	node.ResetStaleCount()
	return node, nil
}

// func NewNodeWithPublicKey(publicKey []byte) (*Node, error) {
// 	id, err := NewID(publicKey)
// 	if err != nil {
// 		err = errors.Wrap(err, "error generating id from public key")
// 		return nil, err
// 	}

// 	return NewNodeWithID(id)
// }

// func NewNodeWithID(id ID) (*Node, error) {
// 	if id == "" {
// 		return nil, errors.New("node-id cannot be empty")
// 	}
// 	return new(Node), nil
// }

func (n *Node) Addr() string {
	return n.addr
}

func (n *Node) ID() ID {
	return n.id
}

func (n *Node) UpdateSeen() {
	n.lastSeen = time.Now().UnixNano() / int64(time.Millisecond)
}

func (n *Node) LastSeen() int64 {
	return n.lastSeen
}

func (n *Node) ResetStaleCount() {
	n.staleCount = 0
}
