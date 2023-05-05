package repository

import (
	"git.cyradar.com/license-manager/backend/internal/dto"
	"git.cyradar.com/license-manager/backend/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetDB() *gorm.DB
	FindOneByName(name string) *models.Product
	FindByID(id string) *models.Product
	FindAll(*dto.PaginationDTO) []models.Product
	CreateOne(creatorID string, productCreateDTO *dto.ProductDTO) string
	Update(id string, productUpdate *dto.ProductDTO) *models.Product
}
