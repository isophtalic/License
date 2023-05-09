package database

import (
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var (
	licenseDatabase *PostgresLicenseProvider
)

type PostgresLicenseProvider struct {
	client *postgresDB.Postgres
	db     *gorm.DB
}

func NewPostgresLicenseProvider(tableName string, db interface{}) *PostgresLicenseProvider {
	clt := db.(*postgresDB.Postgres)
	return &PostgresLicenseProvider{
		client: clt,
		db:     clt.GetDB(),
	}
}

func (repo *PostgresLicenseProvider) GetLicenses(per_page, pg, sort string) (licenses []models.License, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, sort)
	licenses = make([]models.License, 0)
	cursor := database.Scopes(helpers.Paginate(&models.License{}, pagination, database)).Find(&licenses)
	return licenses, pagination.GetPage(), pagination.GetTotalPages(), cursor.Error
}

func (repo *PostgresLicenseProvider) Save(cmd *models.License) error {
	database := repo.db
	result := database.Create(&cmd)
	return result.Error
}

func (repo *PostgresLicenseProvider) UpdateLicense(cmd *models.License) error {
	database := repo.db
	result := database.Save(&cmd)
	return result.Error
}

func (repo *PostgresLicenseProvider) FindByOptions(option string, order interface{}, per_page, pg string) (licenses []models.License, page, totalPage int, err error) {
	database := repo.db
	licenses = make([]models.License, 0)
	pagination := helpers.CreatePagination(per_page, pg, "")
	cursor := database.
		Scopes(helpers.Paginate(&models.License{}, pagination, database)).
		Where(option+" = ?", order).
		Find(&licenses)

	return licenses, pagination.GetPage(), pagination.TotalPages, cursor.Error
}

func (repo *PostgresLicenseProvider) FindByName(name string) (*models.License, error) {
	database := repo.db
	license := new(models.License)
	result := database.Where("name = ?", name).First(&license)
	return license, result.Error
}

func (repo *PostgresLicenseProvider) FindById(id string) (*models.License, error) {
	database := repo.db
	license := new(models.License)
	result := database.Where("license_id = ?", id).First(&license)
	return license, result.Error
}

func (repo *PostgresLicenseProvider) SearchOrFilter(valueSearch string, valueFilter, per_page, pg, sort string) (licenses []models.License, page, totalPage int, err error) {
	database := repo.db
	pagination := helpers.CreatePagination(per_page, pg, sort)
	licenses = make([]models.License, 0)
	if valueFilter != "" {
		return findLicenseWithFilter(database, licenses, pagination, valueSearch, valueFilter)
	}
	cursor := database.
		Where("name LIKE ?", "%"+valueSearch+"%").
		Or("description LIKE ?", "%"+valueSearch+"%").
		Or("tags @> ?", "{"+valueSearch+"}").Find(&licenses)
	tx := cursor.Scopes(helpers.Paginate(&models.License{}, pagination, cursor)).
		Find(&licenses)
	return licenses, pagination.GetPage(), pagination.GetTotalPages(), tx.Error
}

func findLicenseWithFilter(database *gorm.DB, licenses []models.License, pagination *dto.PaginationDTO, valueSearch string, valueFilter string) ([]models.License, int, int, error) {
	cursor := database.Where("name LIKE ?", "%"+valueSearch+"%").
		Or("description LIKE ?", "%"+valueSearch+"%").
		Or("tags @> ?", "{"+valueSearch+"}").
		Or("option = ?", valueFilter)
	tx := cursor.Scopes(helpers.Paginate(&models.License{}, pagination, cursor)).
		Find(&licenses)
	return licenses, pagination.GetPage(), pagination.GetTotalPages(), tx.Error
}
