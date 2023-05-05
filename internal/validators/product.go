package validators

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"

	dto "git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/helpers"
)

func ValidateCreateProduct(productDTO *dto.ProductDTO) {
	productDTO = helpers.TrimStruct(productDTO).(*dto.ProductDTO)
	if *productDTO.Name == "" || *productDTO.Address == "" || *productDTO.Company == "" || *productDTO.Email == "" || *productDTO.Phone == "" || *productDTO.Description == "" {
		customError.Throw(http.StatusUnprocessableEntity, "Missing value in information of product")
	}

	if len(*productDTO.Name) <= 3 {
		customError.Throw(http.StatusUnprocessableEntity, "Company's name must be more than 3 characters.")
	}

	if regexp.MustCompile(`[!@#$%^&*()]`).MatchString(*productDTO.Name) {
		customError.Throw(http.StatusUnprocessableEntity, "Name must not contain any special characters.")
	}

	ValidateEmail(*productDTO.Email)
	ValidatePhoneNumber(*productDTO.Phone)
}

func ValidateUpdateProduct(productUpdate *dto.ProductDTO) {
	fields := reflect.TypeOf(productUpdate).Elem()
	values := reflect.ValueOf(productUpdate).Elem()
	for i := 0; i < fields.NumField(); i++ {
		if values.Field(i).IsNil() {
			continue
		}

		currentValue := strings.TrimSpace(values.Field(i).Elem().String())
		if fields.Field(i).Type.Elem().Kind() == reflect.String && currentValue == "" {
			customError.Throw(http.StatusUnprocessableEntity, "Invalid value: Value contains only space")
		}

		switch strings.ToLower(fields.Field(i).Name) {
		case "company":
			if len(currentValue) <= 3 {
				customError.Throw(http.StatusUnprocessableEntity, "Company's name must be more than 3 characters.")
			}

			if regexp.MustCompile(`[!@#$%^&*()]`).MatchString(currentValue) {
				customError.Throw(http.StatusUnprocessableEntity, "Name must not contain any special characters.")
			}
		case "email":
			ValidateEmail(currentValue)
		case "phone":
			ValidatePhoneNumber(currentValue)
		default:
		}
	}

}
