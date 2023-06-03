package repository

import (
	"github.com/isophtalic/License/internal/models"
	"gorm.io/gorm"
)

type LicenseKeyRepository interface {
	GetDB() *gorm.DB
	Create(keys ...*models.License_key)
	GetByLicenseID(license_id string) ([]models.License_key, error)
	ChangeStatus(license_id string, status bool) error
}
