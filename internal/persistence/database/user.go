package database

import (
	"time"

	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var (
	UserDatabase *PostgresUserProvider
)

type PostgresUserProvider struct {
	client *postgresDB.Postgres
	db     *gorm.DB
}

func NewPostgresUserProvider(tableName string, db interface{}) *PostgresUserProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()
	UserDatabase = &PostgresUserProvider{
		client: db.(*postgresDB.Postgres),
		db:     database,
	}
	return UserDatabase
}

func (repo *PostgresUserProvider) GetUsers(per_page, pg, sort string) (users []models.User, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, sort)
	users = make([]models.User, 0)
	cursor := database.Scopes(helpers.Paginate(&models.User{}, pagination, database)).Find(&users)

	return users, pagination.GetPage(), pagination.GetTotalPages(), cursor.Error
}

func (repo *PostgresUserProvider) Save(user models.User) error {
	database := repo.db
	result := database.Create(&user)
	return result.Error
}

func (repo *PostgresUserProvider) UpdateProfile(cmd models.User) error {
	database := repo.db
	result := database.Save(&cmd)
	return result.Error
}

func (repo *PostgresUserProvider) UpdateLastLogin(user models.User) error {
	database := repo.db
	result := database.Model(&user).Update("last_logged_in", time.Now())
	return result.Error
}

func (repo *PostgresUserProvider) FindByID(id string) (*models.User, error) {
	database := repo.db
	user := new(models.User)
	result := database.Where("user_id = ?", id).First(&user)
	return user, result.Error
}

func (repo *PostgresUserProvider) SearchOrFilter(valueSearch, valueFilter, per_page, pg, sort string) (users []models.User, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, sort)
	users = make([]models.User, 0)

	if valueFilter != "" {
		cursor := database.Where("name LIKE ?", "%"+valueSearch+"%").Or("email LIKE ?", "%"+valueSearch+"%").Or("status = ?", valueFilter).Or("role = ?", valueFilter)
		tx := cursor.Scopes(helpers.Paginate(&models.User{}, pagination, cursor)).
			Find(&users)
		return users, pagination.GetPage(), pagination.GetTotalPages(), tx.Error
	}

	cursor := database.Where("name LIKE ?", "%"+valueSearch+"%").Or("email LIKE ?", "%"+valueSearch+"%")
	tx := cursor.Scopes(helpers.Paginate(&models.User{}, pagination, cursor)).
		Find(&users)
	return users, pagination.GetPage(), pagination.GetTotalPages(), tx.Error

}

func (repo *PostgresUserProvider) SearchByEmail(email string) (*models.User, error) {
	database := repo.db
	user := new(models.User)
	result := database.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (repo *PostgresUserProvider) FilterByOptions(optionsKey string, optionsValue interface{}, per_page, pg string) (users []models.User, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, "")
	users = make([]models.User, 0)

	cursor := database.Where(optionsKey+" = ?", optionsValue)
	tx := cursor.
		Scopes(helpers.Paginate(&models.User{}, pagination, cursor)).
		Find(&users)

	return users, pagination.GetPage(), pagination.TotalPages, tx.Error

}
