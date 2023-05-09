package repository

import (
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence/database"
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
