package main

import (
	"encoding/hex"
	"fmt"
)

type Chain struct {
	Blocks []Block
}

func (c *Chain) FindUTXOs(address string) (int, map[string][]UTXO) {
	spentTXOs := c.FindSpentTXOs(address)
	utxos := make(map[string][]UTXO)
	balance := 0

	for _, block := range c.Blocks {
		for _, tx := range block.Transactions {
			txId := hex.EncodeToString(tx.Id)

			for outIdx, output := range tx.Outputs {
				if output.Receiver == address {
					if spentTXOs[txId] != nil {
						for _, refId := range spentTXOs[txId] {
							if refId != outIdx {
								utxos[txId] = append(utxos[txId],
									UTXO{
										Output:   output,
										OutputId: outIdx,
									},
								)
								balance += output.Amount
							}
						}
					} else {
						utxos[txId] = append(utxos[txId],
							UTXO{
								Output:   output,
								OutputId: outIdx,
							},
						)
						balance += output.Amount
					}
				}
			}
		}
	}

	return balance, utxos
}

func (c *Chain) FindSpentTXOs(address string) map[string][]int {
	spentTXOs := make(map[string][]int)

	for _, block := range c.Blocks {
		for _, tx := range block.Transactions {
			for _, inTx := range tx.Inputs {
				if inTx.PrivKey == address {
					refTxId := hex.EncodeToString(inTx.TxId)
					spentTXOs[refTxId] = append(spentTXOs[refTxId], inTx.OutId)
				}
			}
		}
	}

	return spentTXOs
}

func (c *Chain) Prepend(block Block) {
	if len(c.Blocks) == 0 {
		c.Blocks = append(c.Blocks, block)
	} else {
		c.Blocks = append(c.Blocks, Block{})
		copy(c.Blocks[1:], c.Blocks)
		c.Blocks[0] = block
	}
}

func (c *Chain) Print(addresses []string) {
	for i := len(c.Blocks) - 1; i >= 0; i-- {
		b := c.Blocks[i]
		fmt.Printf("Block #%d\n", len(c.Blocks)-i)
		fmt.Printf("Hash: %s\n", hex.EncodeToString(b.Hash))
		fmt.Printf("Prev hash: %s\n", hex.EncodeToString(b.PrevHash))
		fmt.Printf("Merkle hash: %s\n", hex.EncodeToString(b.MerkleHash))
		fmt.Printf("Transactions:\n")
		for _, tx := range b.Transactions {
			fmt.Printf("\tTx hash: %s\n", hex.EncodeToString(tx.Id))
		}
		fmt.Print("\n")
	}
	for _, address := range addresses {
		balance, _ := c.FindUTXOs(address)
		fmt.Printf("Address %s has a balance of %d\n", address, balance)
	}
	fmt.Print("\n")
}
