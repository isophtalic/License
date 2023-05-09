package service

import (
	"net/http"

	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func UpdateProductOption(optionID string, updatedProduct *dto.ProductOptionDTO) {
	if updatedProduct == nil {
		customError.Throw(http.StatusInternalServerError, "Updated Product is null pointer")
	}
	validators.ValidateUPdateProductOption(updatedProduct)
	persistence.ProductOption().UpdateByID(optionID, updatedProduct)
}

func UpdateOptionDetail(optionID string, updateOptionDetail []dto.OptionDetailDTO) {

}
