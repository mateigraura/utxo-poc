package mock

var adds = []string{
	"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	"35hK24tcLEWcgNA4JxpvbkNkoAcDGqQPsP",
	"34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo",
}

func Commitments() map[string]string {
	return map[string]string{
		"satoshi": adds[0],
		"alice":   adds[1],
		"bob":     adds[2],
	}
}
