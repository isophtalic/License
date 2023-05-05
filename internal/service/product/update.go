package service

import (
	"git.cyradar.com/license-manager/backend/internal/dto"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"
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
