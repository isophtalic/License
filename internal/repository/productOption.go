package repository

import (
	"git.cyradar.com/license-manager/backend/internal/dto"
	"git.cyradar.com/license-manager/backend/internal/models"
	"gorm.io/gorm"
)

type ProductOptionRepository interface {
	GetDB() *gorm.DB
	Create(option *dto.ProductOptionDTO)
	FindByName(name string) *models.ProductOption
	FindByID(id string) *models.ProductOption
	UpdateByID(id string, updatedData *dto.ProductOptionDTO)
	DeleteByID(id string) *models.ProductOption
}
