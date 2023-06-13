package service

import (
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func Encrypt(license_id string) (encodedCipherText, key string) {
	_, priKeyString, data := getKeys(license_id)
	hashedStringInfo, KeyAES, initialVector := hashInfoByAES(&data)

	privateKey := libUtilities.PEMStringToRSAPrivateKey([]byte(priKeyString))
	encodedCipherText = newRSA.Sign(hashedStringInfo, privateKey.(*rsa.PrivateKey))

	// add key AES inside license_key
	encodedCipherText += "." + base64.StdEncoding.EncodeToString(KeyAES) + "." + base64.StdEncoding.EncodeToString(initialVector)

	key = libUtilities.KeyGenerateString()
	bol := true
	licenseKey := &models.License_key{
		Key:       &key,
		LicenseID: &license_id,
		Status:    &bol,
		Start:     time.Now(),
		End:       time.Now().Add(time.Hour * 24 * 365),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	persistence.LicenseKey().Create(licenseKey)
	return
}

func Active(c *gin.Context) error {
	key := c.Param("license_key")
	licenseID := c.Param("license_id")

	if len(key) == 0 || len(licenseID) == 0 {
		return errors.New("Need to license_id")
	}

	status := true

	err := persistence.LicenseKey().ChangeStatus(licenseID, key, status)
	return err
}

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

func hashInfoByAES(info interface{}) (string, []byte, []byte) {
	byteInfo, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	key := newAES.GenerateKeyBYTES(32)
	AESBlock := newAES.MakeCipherBlock(key)
	initialVector := newAES.GenerateInitializationVector()
	hashFromData := newAES.Encrypt(AESBlock, initialVector, string(byteInfo))
	hashStringFromData := base64.StdEncoding.EncodeToString(hashFromData)
	return hashStringFromData, key, initialVector
}

func unHashInfoAES(block cipher.Block, iv []byte, cipherText string) []byte {
	return newAES.Decrypt(block, iv, cipherText)
}
