package zkp

import (
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/r1cs"
	"log"
)

type Prover struct {
	R1CS r1cs.R1CS
	Pk   groth16.ProvingKey
}

func NewProver(r1cs r1cs.R1CS, pk groth16.ProvingKey) Prover {
	return Prover{r1cs, pk}
}

func (p *Prover) Prove(commitment *HashCircuit) groth16.Proof {
	proof, err := groth16.Prove(p.R1CS, p.Pk, commitment)
	if err != nil {
		log.Panic(err)
	}
	return proof
}
