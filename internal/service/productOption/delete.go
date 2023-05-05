package service

import (
	"fmt"
	"net/http"

	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
)

func DeleteOptionByID(id string) *models.ProductOption {
	option := persistence.ProductOption().FindByID(id)
	if option == nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("Option with ID: %v not found.", id))
	}
	deletedOption := persistence.ProductOption().DeleteByID(id)
	return deletedOption
}
