package tornado

import (
	"fmt"
	"github.com/mateigraura/utxo-poc/mock"
	"github.com/mateigraura/utxo-poc/zkp"
)

type Wallet struct {
	Owner     string
	Addresses map[string]string
}

func (w *Wallet) NewWallet(owner string) {
	w.Addresses = mock.Commitments()
	w.Owner = owner
}

// this would be a web3 contract instance. the wallet would only sign the tx
func (w *Wallet) MixerDeposit(receiverAddress string, client Client, ScAddress Mixer) {
	fmt.Printf("Signing TX for x coins for owner: %s\n", w.Owner)
	fmt.Printf("Receiver address: %s\n", receiverAddress)

	var txPayload zkp.HashCircuit
	receiver := w.Addresses[receiverAddress]
	commitment := txPayload.Commit([]byte(receiver))

	// we assume the Mixer is a smart contract at some ScAddress
	client.SendFunds(txPayload, commitment, ScAddress)
}

// this would be a web3 contract instance. the wallet would only sign the tx
func (w *Wallet) MixerWithdraw(client Client, ScAddress Mixer) {
	fmt.Printf("Initiating claim TX for x coins for owner: %s\n", w.Owner)

	var txPayload zkp.HashCircuit
	walletAddress := w.Addresses[w.Owner]
	commitment := txPayload.Commit([]byte(walletAddress))

	client.Withdraw(txPayload, commitment, ScAddress)
}
