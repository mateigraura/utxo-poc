package mock

import (
	"encoding/hex"
	"github.com/mateigraura/utxo-poc/zkp"
)

var adds = []string{
	"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	"35hK24tcLEWcgNA4JxpvbkNkoAcDGqQPsP",
	"34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo",
}

func Commitments() map[string]zkp.CircuitMock {
	var us [][]byte
	for _, add := range adds {
		bh, _ := hex.DecodeString(add)
		us = append(us, bh)
	}

	return map[string]zkp.CircuitMock{
		"satoshi": {
			ReceiverEncoded: us[0],
			Receiver:        adds[0],
			Commitment:      "16452782404320223144130378502757799421584249165328353070394515258282950537369",
		},
		"alice": {
			ReceiverEncoded: us[1],
			Receiver:        adds[1],
			Commitment:      "16604665012224128598511601280298912200197661490959917546275616112006759974009",
		},
		"bob": {
			ReceiverEncoded: us[2],
			Receiver:        adds[2],
			Commitment:      "8013924876658092473206451215532440734759000000000000000000000000000000000000",
		},
	}
}
