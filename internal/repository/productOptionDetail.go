package repository

import (
	"git.cyradar.com/license-manager/backend/internal/dto"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence/database"
	"gorm.io/gorm"
)

type ProductOptionDetailRepository interface {
	GetDB() *gorm.DB
	WithTransaction(tx *gorm.DB) *database.PostgresPdOptionDetailProvider
	Create(productOptionID string, optionDetail ...dto.OptionDetailDTO)
	FindByKey(key string) *models.OptionDetail
	FindByID(id string) *models.OptionDetail
	DeleteByID(id string) *models.OptionDetail
}
