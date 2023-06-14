package service

import (
	"mime/multipart"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func Create(creatorEmail string, productDTO *dto.ProductDTO) string {
	validators.ValidateCreateProduct(productDTO)
	existingProduct := persistence.Product().FindOneByName(*productDTO.Name)
	if existingProduct != nil {
		customError.Throw(http.StatusConflict, "Name of product was already in use.")
	}
	return persistence.Product().CreateOne(creatorEmail, productDTO)
}

func GenerateKeys(productID string, creatorEmail string, typeKey string) {
	creator, err := persistence.User().SearchByEmail(creatorEmail)
	if creator == nil || err != nil {
		customError.Throw(http.StatusNotFound, "Something went wrong while finding user.")
	}

	product := persistence.Product().FindByID(productID)
	if product == nil {
		customError.Throw(http.StatusNotFound, "Something went wrong while finding product.")
	}

	privateKey, publicKey := helpers.MarshalAsPEMStr(helpers.GenerateKeyPair(typeKey))
	key := models.Key{
		KeyID:        &[]string{uuid.NewString()}[0],
		Type:         &typeKey,
		PrivateKey:   &privateKey,
		PublicKey:    &publicKey,
		CreatedAt:    &[]time.Time{time.Now()}[0],
		CreatorID_FK: creator.UserID,
		ProductID_FK: &productID,
	}
	persistence.Key().Create(&key)
}

func UploadKeys(productID string, creatorEmail string, typeKey string, files []*multipart.FileHeader) {
	creator, err := persistence.User().SearchByEmail(creatorEmail)
	if creator == nil || err != nil {
		customError.Throw(http.StatusNotFound, "Something went wrong while finding user.")
	}

	product := persistence.Product().FindByID(productID)
	if product == nil {
		customError.Throw(http.StatusNotFound, "Something went wrong while finding product.")
	}
	key := models.Key{
		KeyID:        &[]string{uuid.NewString()}[0],
		Type:         &typeKey,
		PrivateKey:   &[]string{helpers.ReadFile(files[0])}[0],
		PublicKey:    &[]string{helpers.ReadFile(files[1])}[0],
		CreatedAt:    &[]time.Time{time.Now()}[0],
		CreatorID_FK: creator.UserID,
		ProductID_FK: &productID,
		Creator:      creator,
	}
	persistence.Key().Create(&key)
}
