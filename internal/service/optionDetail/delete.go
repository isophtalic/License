package service

import (
	"fmt"
	"net/http"

	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/persistence"
)

func DeleteByID(id string) {
	opDetail := persistence.ProductOptionDetail().FindByID(id)
	if opDetail == nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("Option detail with ID: '%v' is not found", id))
	}
	persistence.ProductOptionDetail().DeleteByID(id)
}
