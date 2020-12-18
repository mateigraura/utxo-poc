package zkp

import (
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/r1cs"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/io"
	"github.com/consensys/gurvy"
	"log"
)

type Setup struct {
	R1CS r1cs.R1CS
}

func (s *Setup) CompileCircuit(curveId gurvy.ID, circuit *HashCircuit) {
	_r1cs, err := frontend.Compile(curveId, circuit)
	if err != nil {
		log.Panic(err)
	}
	s.R1CS = _r1cs
}

func (s *Setup) ComputeKeys(prover *Prover, verifier *Verifier) {
	if s.R1CS == nil {
		log.Panic("Compiled circuit required!")
	}
	pk, vk := groth16.Setup(s.R1CS)
	prover.R1CS = s.R1CS
	prover.Pk = pk
	verifier.Vk = vk
}

func (s *Setup) Save(path string) {
	err := io.WriteFile(path, s.R1CS)
	if err != nil {
		log.Panic(err)
	}
}
