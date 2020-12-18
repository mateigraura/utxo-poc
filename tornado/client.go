package tornado

import (
	"fmt"
	"github.com/mateigraura/utxo-poc/zkp"
)

type Client struct {
	Prover zkp.Prover
}

func (c *Client) SendFunds(commitment zkp.HashCircuit, nullifier string, m Mixer) {
	err := m.Deposit(commitment, nullifier)
	if err != nil {
		fmt.Printf("ERROR 'deposit': %s\n\n", err)
	} else {
		fmt.Printf("Succesful deposit\n\n")
	}
}

func (c *Client) Withdraw(commitment zkp.HashCircuit, nullifier string, m Mixer) {
	proof := c.Prover.Prove(&commitment)
	err := m.Claim(proof, nullifier)
	if err != nil {
		fmt.Printf("ERROR 'claim': %s\n\n", err)
	} else {
		fmt.Printf("Proof accepted. Coins transfered\n\n")
	}
}
