package main

import (
	"container/list"
	"crypto/sha256"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

type MerkleTree struct {
	Root       *MerkleNode
	MerkleHash []byte
}

func CreateMerkleTree(txs []Tx) *MerkleTree {
	var data [][]byte
	for _, tx := range txs {
		data = append(data, tx.Encode())
	}

	nodes := MakeLeaves(data)

	for nodes.Len() > 1 {
		for i := 0; i < nodes.Len(); i += 2 {
			left := nodes.Front()
			nodes.Remove(left)
			right := nodes.Front()
			nodes.Remove(right)

			nodes.PushBack(MakeNode(
				left.Value.(*MerkleNode),
				right.Value.(*MerkleNode),
			))
		}
	}

	root := nodes.Front().Value.(*MerkleNode)
	return &MerkleTree{root, root.Hash}
}

func MakeLeaves(data [][]byte) *list.List {
	leaves := list.New()

	if len(data)&1 == 1 {
		data = append(data, data[len(data)-1])
	}

	for _, v := range data {
		leaf := MakeLeaf(v)
		leaves.PushBack(leaf)
	}

	return leaves
}

func MakeNode(left *MerkleNode, right *MerkleNode) *MerkleNode {
	node := MerkleNode{}

	node.Left = left
	node.Right = right

	joinedHash := append(left.Hash, right.Hash...)
	leavesHash := sha256.Sum256(joinedHash)
	node.Hash = leavesHash[:]

	return &node
}

func MakeLeaf(data []byte) *MerkleNode {
	node := MerkleNode{}
	node.Left = nil
	node.Right = nil

	hash := sha256.Sum256(data)
	node.Hash = hash[:]

	return &node
}
