package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"

	customError "github.com/isophtalic/License/internal/error"
)

func GenerateKeyPair(keyType string) (*rsa.PrivateKey, *rsa.PublicKey) {
	switch keyType {
	case "RSA2048":
		return generateKeysRSA2048()
	default:
		customError.Throw(http.StatusNotImplemented, fmt.Sprintf("%v encryption does not support yet.", keyType))
	}
	return nil, nil
}

func generateKeysRSA2048() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}

func MarshalAsPEMStr(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, string) {
	return MarshalPrivKeyAsPEMStr(privateKey), MarshalPubKeyAsPEMStr(publicKey)
}

func MarshalPubKeyAsPEMStr(publicKey *rsa.PublicKey) string {
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC-KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	))
	return pubKeyPem
}

func MarshalPrivKeyAsPEMStr(privateKey *rsa.PrivateKey) string {
	privKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE-KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	))
	return privKeyPem

}
