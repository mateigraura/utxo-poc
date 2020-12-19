package tornado

import (
	"github.com/consensys/gurvy"
	"github.com/mateigraura/utxo-poc/zkp"
)

func Run() {
	var mixer Mixer
	var client Client
	var setup zkp.Setup
	var mimcCirc zkp.HashCircuit

	mixer.Init()

	// init constraint system & generate pk, vk
	setup.CompileCircuit(gurvy.BN256, &mimcCirc)
	setup.ComputeKeys(&client.Prover, &mixer.Verifier)

	var satoshiWallet, aliceWallet, bobWallet Wallet

	satoshiWallet.NewWallet("satoshi")
	aliceWallet.NewWallet("alice")
	bobWallet.NewWallet("bob")

	// satoshi deposits x coins for alice
	satoshiWallet.MixerDeposit("alice", client, mixer)

	// alice withdraws x coins received from satoshi
	aliceWallet.MixerWithdraw(client, mixer)

	// if alice tries to withdraw again it will fail
	aliceWallet.MixerWithdraw(client, mixer)

	// satoshi deposits x coins for bob
	satoshiWallet.MixerDeposit("bob", client, mixer)

	// bob withdraws x coins received from satoshi
	bobWallet.MixerWithdraw(client, mixer)
}
