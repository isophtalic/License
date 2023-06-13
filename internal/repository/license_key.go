package repository

import (
	"github.com/isophtalic/License/internal/models"
	"gorm.io/gorm"
)

type LicenseKeyRepository interface {
	GetDB() *gorm.DB
	Create(keys ...*models.License_key)
	GetLicenseKey(per_page, pg, sort string, license_id string) (licenseKeys []models.License_key, page, totalPage int, err error)
	GetByLicenseID(license_id string) ([]models.License_key, error)
	ChangeStatus(license_id string, key string, status bool) error
}
