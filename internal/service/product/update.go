package service

import (
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func Update(id string, productUpdate *dto.ProductDTO) *models.Product {
	validators.ValidateUpdateProduct(productUpdate)
	updatedProduct := persistence.Product().Update(id, productUpdate)
	return updatedProduct
}

func ChangeStatus(id string, status bool) {
	persistence.Product().Update(id, &dto.ProductDTO{
		Status: &status,
	})

}
