package zkp

import "github.com/consensys/gnark/backend/groth16"

type Verifier struct {
	Vk groth16.VerifyingKey
}

func NewVerifier(vk groth16.VerifyingKey) Verifier {
	return Verifier{vk}
}

func (v *Verifier) Verify(proof groth16.Proof, commitment *HashCircuit) (bool, error) {
	err := groth16.Verify(proof, v.Vk, commitment)
	if err != nil {
		return false, err
	}
	return true, nil
}
