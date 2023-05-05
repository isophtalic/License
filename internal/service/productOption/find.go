package service

import (
	"fmt"
	"net/http"

	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"gorm.io/gorm"
)

func DetailProductOption(id string) *models.ProductOption {
	var option models.ProductOption
	r := persistence.ProductOption().GetDB().Preload("OptionDetails", func(db *gorm.DB) *gorm.DB {
		return db.Select("option_detail_id", "key", "value", "created_at", "product_option_id")
	}).Select("option_id", "name", "description", "enable").First(&option)
	if r.Error != nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("option with ID:'%v' was not found.", id))
	}
	return &option
}
