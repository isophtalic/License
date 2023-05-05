package repository

import "git.cyradar.com/license-manager/backend/internal/models"

type LicenseRepository interface {
	GetLicenses(per_page, pg, sort string) (licenses []models.License, page, totalPage int, err error)
	Save(cmd *models.License) error
	UpdateLicense(cmd *models.License) error
	FindByName(name string) (*models.License, error)
	FindById(id string) (*models.License, error)
	FindByOptions(option string, order interface{}, per_page, pg string) (licenses []models.License, page, totalPage int, err error)
	SearchOrFilter(valueSearch string, valueFilter, per_page, pg, sort string) (licenses []models.License, page, totalPage int, err error)
}
