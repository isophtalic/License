package validators

import (
	"net/http"

	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/helpers"
)

func ValidateProductOptionDetail(pdOptionDetail *dto.OptionDetailDTO) {
	if pdOptionDetail == nil {
		customError.Throw(http.StatusBadRequest, "Product-options-detail must not be null")
	}

	pdOptionDetail = helpers.TrimStruct(pdOptionDetail).(*dto.OptionDetailDTO)
	println("OKDEOKO", *pdOptionDetail.Key == "", *pdOptionDetail.Value == "")
	if *pdOptionDetail.Key == "" || *pdOptionDetail.Value == "" {
		customError.Throw(http.StatusUnprocessableEntity, "Missing value of product option detail.")
	}
}

// func ValidateUpdateOptionDetail(option dto.OptionDetailDTO) {
// 	if strings.TrimSpace(*option.Key) == "" || strings.TrimSpace(*option.Value) == "" {
// 		customError.Throw(http.StatusUnprocessableEntity, "Invalid: Key and Value option must be not empty")
// 	}
// }
