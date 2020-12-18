package zkp

import (
	"encoding/hex"
	"github.com/consensys/gnark/crypto/hash/mimc/bn256"
	"github.com/consensys/gurvy/bn256/fr"
)

func MiMCBn256Decoded(sc string) string {
	scBytes := []byte(sc)
	mimc := MiMCBn256(scBytes)
	decoded := hex.EncodeToString(mimc)
	return decoded
}

func MiMCBn256(sc []byte) []byte {
	var data, tmp fr.Element
	data.SetBytes(sc)
	dataBytes := data.Bytes()
	b := bn256.Sum(seed, dataBytes[:])
	tmp.SetBytes(b)
	return tmp.Bytes()
}
