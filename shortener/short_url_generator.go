package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha254Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58Encoded(byte []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(byte)

	if err != nil {
		log.Panicf("Failed to encode %v", err)
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortUrl(initialLink, userId string) string {
	urlHashByte := sha254Of(initialLink + userId)

	generatedNum := new(big.Int).SetBytes(urlHashByte).Uint64()

	finalStr := base58Encoded([]byte(
		fmt.Sprintf("%d", generatedNum),
	))

	return finalStr[:8]
}
