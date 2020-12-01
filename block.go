package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Block struct {
	Hash         []byte
	PrevHash     []byte
	MerkleHash   []byte
	MerkleRoot   *MerkleNode
	Transactions []Tx
}

func NewBlock(txs []Tx, prevHash []byte) Block {
	merkleTree := CreateMerkleTree(txs)

	block := Block{
		Hash:         []byte{},
		PrevHash:     prevHash,
		MerkleHash:   merkleTree.MerkleHash,
		MerkleRoot:   merkleTree.Root,
		Transactions: txs,
	}
	HashBlock(&block)

	return block
}

func GenesisBlock(to string) Block {
	return NewBlock([]Tx{CoinbaseTx(to)}, []byte{})
}

func HashBlock(block *Block) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(block)
	hash := sha256.Sum256(buffer.Bytes())
	block.Hash = hash[:]
}
