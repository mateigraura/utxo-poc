package zkp

import (
	"encoding/hex"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gurvy"
)

const seed = "testSeed"

type HashCircuit struct {
	Receiver   frontend.Variable
	Commitment frontend.Variable `gnark:",public"`
}

type CircuitMock struct {
	ReceiverEncoded []byte
	Receiver        string
	Commitment      string
}

func (hc *HashCircuit) Define(curveID gurvy.ID, cs *frontend.ConstraintSystem) error {
	mimcFunc, err := mimc.NewMiMC(seed, curveID)
	if err != nil {
		return err
	}
	cs.AssertIsEqual(hc.Commitment, mimcFunc.Hash(cs, hc.Receiver))
	return nil
}

func (hc *HashCircuit) Commit(sc []byte) {
	mimcHash := MiMCBn256(sc)
	hc.Commitment.Assign(mimcHash)
	hc.Receiver.Assign(hex.EncodeToString(sc))
}
