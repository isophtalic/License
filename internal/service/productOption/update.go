package service

import (
	"net/http"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"
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
