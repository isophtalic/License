package validators

import (
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/helpers"
	"net/http"
)

func ValidateProductOption(option *dto.ProductOptionDTO) {
	if option == nil {
		customError.Throw(http.StatusUnprocessableEntity, "Validate: Option is null pointer")
	}
	option = helpers.TrimStruct(option).(*dto.ProductOptionDTO)
	if *option.Name == "" || *option.Description == "" {
		customError.Throw(http.StatusUnprocessableEntity, "Missing value in information of product option")
	}
}

func ValidateUPdateProductOption(option *dto.ProductOptionDTO) {
	if option == nil {
		customError.Throw(http.StatusUnprocessableEntity, "Validate: Option is null pointer")
	}
	option = helpers.TrimStruct(option).(*dto.ProductOptionDTO)
	if option.Name != nil && *option.Name == "" {
		customError.Throw(http.StatusUnprocessableEntity, "Validate: name must not be empty")
	}

	if option.Description != nil && *option.Description == "" {
		customError.Throw(http.StatusUnprocessableEntity, "Validate: Description must not be empty")
	}
}
