package database

import (
	"fmt"

	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var LicenseKeyDatabase *PostgresLicenseKeyProvider

type PostgresLicenseKeyProvider struct {
	table string
	db    *gorm.DB
}

func NewPostgresLicenseKeyProvider(tableName string, db interface{}) *PostgresLicenseKeyProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()

	LicenseKeyDatabase = &PostgresLicenseKeyProvider{
		table: tableName,
		db:    database,
	}

	return LicenseKeyDatabase
}

func (repo *PostgresLicenseKeyProvider) GetDB() *gorm.DB {
	return repo.db
}

func (repo *PostgresLicenseKeyProvider) Create(keys ...*models.License_key) {
	if len(keys) == 0 {
		return
	}
	r := repo.db.Model(&models.License_key{}).Create(keys)
	if r.Error != nil {
		println(fmt.Printf("Error: Create keys:::%v", r.Error))
		panic(r.Error)
	}
}

func (repo *PostgresLicenseKeyProvider) GetByLicenseID(license_id string) ([]models.License_key, error) {
	database := repo.db
	license_key := make([]models.License_key, 0)
	result := database.Where("product_id", license_id).Find(&license_key)
	return license_key, result.Error
}

func (repo *PostgresLicenseKeyProvider) ChangeStatus(licene_id string, status bool) error {
	database := repo.db
	result := database.Model(&models.License_key{}).Update("status", &status)
	return result.Error
}
