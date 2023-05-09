package repository

import "github.com/isophtalic/License/internal/models"

type LicenseConfigRepository interface {
	GetLicenseConfigs(per_page, pg, sort string) (configs []models.LicenseConfig, page, totalPage int, err error)
	GetConfigsByLicenseId(per_page, pg, sort, license_id string) (licenseConfigs []models.LicenseConfig, page, totalPage int, err error)
	FindByName(name string) (licenseConfig *models.LicenseConfig, err error)
	FindByID(id string) (licenseConfig *models.LicenseConfig, err error)
	Search(value string, license_id string, per_page, pg, sort string) (licenseConfigs []models.LicenseConfig, page, totalPage int, err error)
	Save(cmd *models.LicenseConfig) error
	UpdateLicense(cmd *models.LicenseConfig) error
}
