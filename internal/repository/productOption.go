package repository

import (
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
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
