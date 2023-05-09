package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func CreateLicense(cmd *dto.CreateLicenseDTO) {
	validators.ValidateCreateLicenseStruct(cmd)
	_, err := persistence.License().FindByName(cmd.Name)
	if err == nil && err.Error() == "record not found" {
		customError.Throw(http.StatusFound, "Name was existed")
		return
	}
	if err == nil && err.Error() != "record not found" {
		customError.Throw(http.StatusFound, err.Error())
		return
	}

	customer, product := validateCustomerAndProduct(cmd)

	id := uuid.NewString()
	expireDate := validators.ValidateExpireDate(cmd.ExpiryDate)
	newLicense := &models.License{
		LicenseID:   &id,
		Name:        &cmd.Name,
		Option:      &cmd.Option,
		Description: &cmd.Description,
		Tag:         cmd.Tag,
		ExpiryDate:  expireDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CustomerID:  customer.CustomerID,
		ProductID:   product.ProductID,
	}
	err = persistence.License().Save(newLicense)
	if err != nil {
		customError.Throw(http.StatusInternalServerError, err.Error())
	}
}

func validateCustomerAndProduct(cmd *dto.CreateLicenseDTO) (*models.Customer, *models.Product) {
	customer, err := persistence.Customer().FindByName(cmd.CustomerName)
	if err != nil {
		fmt.Println("err: ", err.Error())
		customError.Throw(http.StatusUnprocessableEntity, "Customer is invalid")
		return nil, nil
	}
	if !*customer.Enable {
		customError.Throw(http.StatusUnprocessableEntity, "Customer is disable")
		return nil, nil
	}

	product := persistence.Product().FindOneByName(cmd.ProductName)
	if !*product.Status {
		fmt.Println(*product.Status)
		customError.Throw(http.StatusUnprocessableEntity, "Product is disable")
		return nil, nil
	}
	return customer, product
}
