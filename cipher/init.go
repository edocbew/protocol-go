package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
)

const (
	keySize = 2048
)

func GenerateKeyPairs() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, errPrivKey := rsa.GenerateKey(rand.Reader, keySize)
	if errPrivKey != nil {
		fmt.Println("Error in generating Private Key!")
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}
