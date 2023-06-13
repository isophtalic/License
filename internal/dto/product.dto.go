package dto

import (
	"time"

	"github.com/isophtalic/License/internal/models"
)

type ProductDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Company     *string `json:"company,omitempty"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Address     *string `json:"address,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

func ToProduct(productDTO *ProductDTO) *models.Product {
	return &models.Product{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Company:     productDTO.Company,
		Email:       productDTO.Email,
		Phone:       productDTO.Phone,
		Address:     productDTO.Address,
		Status:      productDTO.Status,
	}
}

func UpdateProduct(productDTO *ProductDTO) *models.Product {
	return &models.Product{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Company:     productDTO.Company,
		Email:       productDTO.Email,
		Phone:       productDTO.Phone,
		Address:     productDTO.Address,
		Status:      productDTO.Status,
		UpdatedAt:   time.Now(),
	}
}
