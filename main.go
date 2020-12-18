package main

import (
	"github.com/mateigraura/utxo-poc/chain"
	"github.com/mateigraura/utxo-poc/tornado"
	"os"
)

func main() {

	opt := os.Args[1]
	switch opt {
	case "utxo-tx":
		chain.Run()
		break
	case "tornado":
		tornado.Run()
		break
	}
}
