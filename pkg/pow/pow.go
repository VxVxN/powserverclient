package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

const (
	leadingZeros = 4
)

func PerformPoW(challenge string) string {
	var nonce int
	for {
		data := challenge + " " + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(data))
		if hasLeadingZeros(hash[:]) {
			return strconv.Itoa(nonce)
		}
		nonce++
	}
}

func hasLeadingZeros(hash []byte) bool {
	hashHex := hex.EncodeToString(hash)
	return strings.HasPrefix(hashHex, strings.Repeat("0", leadingZeros))
}
