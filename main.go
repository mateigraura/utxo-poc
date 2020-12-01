package main

func main() {
	satoshi := "12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX"
	alice := "35hK24tcLEWcgNA4JxpvbkNkoAcDGqQPsP"
	bob := "34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo"
	addresses := []string{satoshi, alice, bob}

	chain := Chain{}
	chain.Prepend(GenesisBlock(satoshi))

	chain.Print(addresses)

	tx1 := NewTx(satoshi, alice, 15, chain)

	block1 := NewBlock([]Tx{*tx1, CoinbaseTx("")}, chain.Blocks[0].Hash)

	chain.Prepend(block1)

	chain.Print(addresses)

	tx2 := NewTx(alice, bob, 10, chain)

	block2 := NewBlock([]Tx{CoinbaseTx(""), *tx2}, chain.Blocks[0].Hash)

	chain.Prepend(block2)

	chain.Print(addresses)
}
