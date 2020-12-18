package tornado

import (
	"fmt"
	"github.com/mateigraura/utxo-poc/mock"
	"github.com/mateigraura/utxo-poc/zkp"
)

type Wallet struct {
	Owner string
	Cts   map[string]zkp.CircuitMock
}

func (w *Wallet) NewWallet(owner string) {
	w.Cts = mock.Commitments()
	w.Owner = owner
}

// this would be a web3 contract instance. the wallet would only sign the tx
func (w *Wallet) MixerDeposit(receiverAddress string, client Client, ScAddress Mixer) {
	fmt.Printf("Signing TX for x coins for owner: %s\n", w.Owner)
	fmt.Printf("Receiver address: %s\n", receiverAddress)

	var txPayload zkp.HashCircuit
	receiver := w.Cts[receiverAddress]

	// private info, won't be disclosed
	txPayload.Receiver.Assign(receiver.ReceiverEncoded)
	// public info, usesless without proof
	txPayload.Commitment.Assign(receiver.Commitment)

	// we assume the Mixer is a smart contract at some ScAddress
	client.SendFunds(txPayload, receiver.Commitment, ScAddress)
}

// this would be a web3 contract instance. the wallet would only sign the tx
func (w *Wallet) MixerWithdraw(client Client, ScAddress Mixer) {
	fmt.Printf("Initiating claim TX for x coins for owner: %s\n", w.Owner)

	var txPayload zkp.HashCircuit
	walletAddress := w.Cts[w.Owner]

	// private info, won't be disclosed
	txPayload.Receiver.Assign(walletAddress.ReceiverEncoded)
	// public info, usesless without proof
	txPayload.Commitment.Assign(walletAddress.Commitment)

	client.Withdraw(txPayload, walletAddress.Commitment, ScAddress)
}
