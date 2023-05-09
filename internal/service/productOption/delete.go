package service

import (
	"fmt"
	"net/http"

	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
)

func DeleteOptionByID(id string) *models.ProductOption {
	option := persistence.ProductOption().FindByID(id)
	if option == nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("Option with ID: %v not found.", id))
	}
	deletedOption := persistence.ProductOption().DeleteByID(id)
	return deletedOption
}
