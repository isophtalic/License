package database

import (
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var (
	licenseConfigDatabase *PostgresLicenseConfigProvider
)

type PostgresLicenseConfigProvider struct {
	client *postgresDB.Postgres
	db     *gorm.DB
}

func NewPostgresLicenseConfigProvider(tableName string, db interface{}) *PostgresLicenseConfigProvider {
	clt := db.(*postgresDB.Postgres)
	return &PostgresLicenseConfigProvider{
		client: clt,
		db:     clt.GetDB(),
	}
}

func (repo *PostgresLicenseConfigProvider) GetLicenseConfigs(per_page, pg, sort string) (configs []models.LicenseConfig, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, sort)
	configs = make([]models.LicenseConfig, 0)
	cursor := database.Scopes(helpers.Paginate(&models.License{}, pagination, database)).Find(&configs)

	return configs, pagination.GetPage(), pagination.GetTotalPages(), cursor.Error
}

func (repo *PostgresLicenseConfigProvider) GetConfigsByLicenseId(per_page, pg, sort, license_id string) (licenseConfigs []models.LicenseConfig, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, sort)
	licenseConfigs = make([]models.LicenseConfig, 0)
	cursor := database.Where("license_id = ?", license_id).Scopes(helpers.Paginate(&models.LicenseConfig{}, pagination, database)).Find(&licenseConfigs)

	return licenseConfigs, pagination.GetPage(), pagination.GetTotalPages(), cursor.Error
}

func (repo *PostgresLicenseConfigProvider) FindByName(name string) (licenseConfig *models.LicenseConfig, err error) {
	database := repo.db

	licenseConfig = new(models.LicenseConfig)
	cursor := database.Where("name = ?", name).Find(&licenseConfig)

	return licenseConfig, cursor.Error
}
func (repo *PostgresLicenseConfigProvider) FindByID(id string) (licenseConfig *models.LicenseConfig, err error) {
	database := repo.db

	licenseConfig = new(models.LicenseConfig)
	cursor := database.Where("config_id = ?", id).Find(&licenseConfig)

	return licenseConfig, cursor.Error
}

func (repo *PostgresLicenseConfigProvider) Search(value string, license_id string, per_page, pg, sort string) (licenseConfigs []models.LicenseConfig, page, totalPage int, err error) {
	database := repo.db
	pagination := helpers.CreatePagination(per_page, pg, sort)
	licenseConfigs = make([]models.LicenseConfig, 0)

	if license_id != "" {
		cursor := database.Model(&models.LicenseConfig{LicenseID: &license_id}).Where("name LIKE ?", "%"+value+"%").
			Or("value LIKE ?", "%"+value+"%")

		tx := cursor.Scopes(helpers.Paginate(&models.LicenseConfig{}, pagination, cursor)).
			Find(&licenseConfigs)
		return licenseConfigs, pagination.GetPage(), pagination.GetTotalPages(), tx.Error
	}

	cursor := database.Model(&models.LicenseConfig{}).Where("name LIKE ?", "%"+value+"%").
		Or("value LIKE ?", "%"+value+"%")

	tx := cursor.Scopes(helpers.Paginate(&models.LicenseConfig{}, pagination, cursor)).
		Find(&licenseConfigs)
	return licenseConfigs, pagination.GetPage(), pagination.GetTotalPages(), tx.Error

}

func (repo *PostgresLicenseConfigProvider) Save(cmd *models.LicenseConfig) error {
	database := repo.db
	result := database.Create(&cmd)
	return result.Error
}

func (repo *PostgresLicenseConfigProvider) UpdateLicense(cmd *models.LicenseConfig) error {
	database := repo.db
	result := database.Save(&cmd)
	return result.Error
}
