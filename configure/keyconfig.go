package configure

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"
)

type Keys struct {
	DestPubKey []byte
	SrcPrivKey []byte
}

func (k Keys) getSourcePrivateKey() (r *rsa.PrivateKey) {
	r, err := x509.ParsePKCS1PrivateKey(k.SrcPrivKey)
	if err != nil {
		fmt.Println("Error in parsing Private Key to RSA Format!")
		os.Exit(1)
	}
	return
}

func (k Keys) getDestPublicKey() (r *rsa.PublicKey) {
	r, err := x509.ParsePKCS1PublicKey(k.DestPubKey)
	if err != nil {
		fmt.Println("Error in parsing Public Key to RSA Format!")
		os.Exit(1)
	}
	return
}
