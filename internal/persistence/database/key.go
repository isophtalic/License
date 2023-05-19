package database

import (
	"fmt"

	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var KeyDatabase *PostgresKeyProvider

type PostgresKeyProvider struct {
	table string
	db    *gorm.DB
}

func NewPostgresKeyProvider(tableName string, db interface{}) *PostgresKeyProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()

	KeyDatabase = &PostgresKeyProvider{
		table: tableName,
		db:    database,
	}

	return KeyDatabase
}

func (repo *PostgresKeyProvider) GetDB() *gorm.DB {
	return repo.db
}

func (repo *PostgresKeyProvider) Create(keys ...*models.Key) {
	if len(keys) == 0 {
		return
	}
	r := repo.db.Model(&models.Key{}).Create(keys)
	if r.Error != nil {
		println(fmt.Printf("Error: Create keys:::%v", r.Error))
		panic(r.Error)
	}
}

func (repo *PostgresKeyProvider) GetKeyByProductID(productID string) (*models.Key, error) {
	database := repo.db
	key := new(models.Key)
	result := database.Where("product_id", productID).First(&key)
	return key, result.Error
}
