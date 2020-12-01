package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"math/rand"
	"time"
)

type Tx struct {
	Id      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxInput struct {
	TxId    []byte
	OutId   int
	PrivKey string
}

type TxOutput struct {
	Amount   int
	Receiver string
}

type UTXO struct {
	Output   TxOutput
	OutputId int
}

func NewTx(from string, to string, amount int, c Chain) *Tx {
	balance, utxos := c.FindUTXOs(from)

	if balance < amount {
		log.Panic("Insufficient funds")
	}

	var inputs []TxInput
	var outputs []TxOutput

	outputs = append(outputs, TxOutput{Amount: amount, Receiver: to})
	outputs = append(outputs, TxOutput{Amount: balance - amount, Receiver: from})

	for txId, outputs := range utxos {
		txHash, _ := hex.DecodeString(txId)

		for _, output := range outputs {
			input := TxInput{TxId: txHash, OutId: output.OutputId, PrivKey: from}
			inputs = append(inputs, input)
		}
	}

	tx := Tx{Id: nil, Inputs: inputs, Outputs: outputs}
	HashTx(&tx)

	return &tx
}

func CoinbaseTx(to string) Tx {
	rand.Seed(time.Now().UnixNano())
	reward := rand.Intn(40 - 20) + 20
	input := TxInput{TxId: nil, OutId: -1, PrivKey: ""}
	output := TxOutput{Amount: reward, Receiver: to}

	tx := Tx{Id: nil, Inputs: []TxInput{input}, Outputs: []TxOutput{output}}
	HashTx(&tx)

	return tx
}

func (tx *Tx) Encode() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(tx)
	return buffer.Bytes()
}

func HashTx(tx *Tx) {
	txHash := sha256.Sum256(tx.Encode())
	tx.Id = txHash[:]
}
