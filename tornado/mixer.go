package tornado

import (
	"errors"
	"fmt"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/mateigraura/utxo-poc/zkp"
)

type Mixer struct {
	Nullifiers  map[string]bool
	Commitments map[string]zkp.HashCircuit
	Verifier    zkp.Verifier
	MerkleTree  *Tree
}

func (m *Mixer) Init() {
	m.Commitments = make(map[string]zkp.HashCircuit)
	m.Nullifiers = make(map[string]bool)
	m.MerkleTree = NewTree()
}

func (m *Mixer) Deposit(commitment zkp.HashCircuit, nullifier string) error {
	if m.Commitments[nullifier] != (zkp.HashCircuit{}) {
		return errors.New("commitment already used")
	}
	err := m.MerkleTree.Insert(nullifier)
	if err != nil {
		return err
	}
	m.Commitments[nullifier] = commitment
	return nil
}

func (m *Mixer) Claim(proof groth16.Proof, nullifier string) error {
	if m.Nullifiers[nullifier] {
		return errors.New("note already claimed")
	}

	var commitment zkp.HashCircuit
	commitment = m.Commitments[nullifier]
	res, err := m.Verifier.Verify(proof, &commitment)
	if !res {
		return err
	}
	if !m.MerkleTree.HasValidRoot(nullifier) {
		return errors.New("invalid merkle root")
	}

	m.Nullifiers[nullifier] = true

	// transfer logic ... esdt.transfer(...)
	fmt.Println("Proof verified. Transfering coins")
	return nil
}
