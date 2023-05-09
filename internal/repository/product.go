package repository

import (
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
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
