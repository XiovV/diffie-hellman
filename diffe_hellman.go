package dh

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

type ECDH struct {
	curve elliptic.Curve
}

type ellipticPublicKey struct {
	X, Y *big.Int
}

type ellipticPrivateKey struct {
	privateKey []byte
}

func NewECDH(curve elliptic.Curve) *ECDH {
	return &ECDH{curve: curve}
}

func (e *ECDH) GenerateKeyPair() (crypto.PublicKey, crypto.PrivateKey) {
	var privateKey crypto.PrivateKey
	var publicKey crypto.PublicKey

	priv, x, y, err := elliptic.GenerateKey(e.curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
	}

	privateKey = &ellipticPrivateKey{privateKey: priv}

	publicKey = &ellipticPublicKey{
		X: x,
		Y: y,
	}

	return publicKey, privateKey
}

func (e *ECDH) GenerateSharedSecret(privKey crypto.PrivateKey, pubKey crypto.PublicKey) ([]byte, error) {
	priv := privKey.(*ellipticPrivateKey)
	pub := pubKey.(*ellipticPublicKey)

	x, _ := e.curve.ScalarMult(pub.X, pub.Y, priv.privateKey)
	return x.Bytes(), nil
}



