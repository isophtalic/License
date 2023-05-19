package service

import (
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"net/http"

	libAES "github.com/isophtalic/GenerateKey/lib/aes"
	libRSA "github.com/isophtalic/GenerateKey/lib/rsa"
	libUtilities "github.com/isophtalic/GenerateKey/utilities"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/persistence"
	serviceLicense "github.com/isophtalic/License/internal/service/license"

	"github.com/isophtalic/License/internal/models"
)

var newRSA libRSA.RSA
var newAES libAES.AES

func getKeys(license_id string) (string, string, *models.License) {
	data, err := serviceLicense.GetLicenseByID(license_id)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, err.Error())
	}

	keyDetail, err := persistence.Key().GetKeyByProductID(*data.ProductID)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, "Don't get product key by :"+err.Error())
	}
	return *keyDetail.PublicKey, *keyDetail.PrivateKey, data
}

func Encrypt(license_id string) (encodedCipherText, key string) {
	_, priKeyString, data := getKeys(license_id)
	hashedStringInfo, _, _ := hashInfoByAES(&data)

	privateKey := libUtilities.PEMStringToRSAPrivateKey([]byte(priKeyString))
	encodedCipherText = newRSA.Sign(hashedStringInfo, privateKey.(*rsa.PrivateKey))
	key = libUtilities.KeyGenerateString()
	return
}

func hashInfoByAES(info interface{}) (string, cipher.Block, []byte) {
	byteInfo, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	key := newAES.GenerateKeyBYTES(32)
	AESBlock := newAES.MakeCipherBlock(key)
	initialVector := newAES.GenerateInitializationVector()
	hashFromData := newAES.Encrypt(AESBlock, initialVector, string(byteInfo))
	hashStringFromData := base64.StdEncoding.EncodeToString(hashFromData)
	return hashStringFromData, AESBlock, initialVector
}

func unHashInfoAES(block cipher.Block, iv []byte, cipherText string) []byte {
	return newAES.Decrypt(block, iv, cipherText)
}
