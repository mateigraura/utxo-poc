package zkp

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gurvy"
	"github.com/consensys/gurvy/bn256/fr"
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

func (hc *HashCircuit) Commit(sc []byte) string {
	var data fr.Element
	data.SetBytes(sc)
	mimcHash := MiMCBn256(data.Bytes())

	// private info, won't be disclosed
	hc.Receiver.Assign(data)
	// public info, usesless without proof
	hc.Commitment.Assign(mimcHash)

	return string(mimcHash)
}
