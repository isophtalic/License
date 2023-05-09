package database

import (
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var (
	CustomerDatabase *PostgresCustomerProvider
)

type PostgresCustomerProvider struct {
	client *postgresDB.Postgres
	db     *gorm.DB
}

func NewPostgresCustomerProvider(tableName string, db interface{}) *PostgresCustomerProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()
	CustomerDatabase = &PostgresCustomerProvider{
		client: db.(*postgresDB.Postgres),
		db:     database,
	}
	// CustomerDatabase.db.Table(tableName)
	return CustomerDatabase
}

func (repo *PostgresCustomerProvider) GetCustomers(per_page, pg, sort string) (customers []models.Customer, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, sort)
	customers = make([]models.Customer, 0)
	cursor := database.Scopes(helpers.Paginate(&models.Customer{}, pagination, database)).Find(&customers)

	return customers, pagination.GetPage(), pagination.GetTotalPages(), cursor.Error
}

func (repo *PostgresCustomerProvider) Save(customer models.Customer) error {
	database := repo.db
	result := database.Create(&customer)
	return result.Error
}

func (repo *PostgresCustomerProvider) UpdateCustomer(cmd models.Customer) error {
	database := repo.db
	result := database.Save(&cmd)
	return result.Error
}

func (repo *PostgresCustomerProvider) SearchOrFilter(valueSearch string, valueFilter, per_page, pg, sort string) (customers []models.Customer, page, totalPage int, err error) {
	database := repo.db
	pagination := helpers.CreatePagination(per_page, pg, sort)
	customers = make([]models.Customer, 0)
	if valueFilter != "" {
		return findCustomerWithFilter(database, customers, pagination, valueSearch, valueFilter)
	}

	cursor := database.Where("name LIKE ?", "%"+valueSearch+"%").
		Or("email LIKE ?", "%"+valueSearch+"%").
		Or("phone_number LIKE ?", "%"+valueSearch+"%").
		Or("address LIKE ?", "%"+valueSearch+"%").
		Or("description LIKE ?", "%"+valueSearch+"%").
		Or("tags @> ?", "{"+valueSearch+"}")

	tx := cursor.Scopes(helpers.Paginate(&models.Customer{}, pagination, cursor)).
		Find(&customers)
	return customers, pagination.GetPage(), pagination.GetTotalPages(), tx.Error

}

func findCustomerWithFilter(database *gorm.DB, customers []models.Customer, pagination *dto.PaginationDTO, valueSearch string, valueFilter string) ([]models.Customer, int, int, error) {
	cursor := database.Where("name LIKE ?", "%"+valueSearch+"%").
		Or("email LIKE ?", "%"+valueSearch+"%").
		Or("phone_number LIKE ?", "%"+valueSearch+"%").
		Or("address LIKE ?", "%"+valueSearch+"%").
		Or("description LIKE ?", "%"+valueSearch+"%").
		Or("tags @> ?", "{"+valueSearch+"}").
		Or("organization = ?", valueFilter).
		Or("enable = ?", valueFilter)
	tx := cursor.Scopes(helpers.Paginate(&models.Customer{}, pagination, cursor)).
		Find(&customers)
	return customers, pagination.GetPage(), pagination.GetTotalPages(), tx.Error
}

func (repo *PostgresCustomerProvider) FindByEmail(email string) (*models.Customer, error) {
	database := repo.db
	customer := new(models.Customer)
	result := database.Where("email = ?", email).First(&customer)
	return customer, result.Error
}

func (repo *PostgresCustomerProvider) FindByName(name string) (*models.Customer, error) {
	database := repo.db
	customer := new(models.Customer)
	result := database.Where("name = ?", name).First(&customer)
	return customer, result.Error
}

func (repo *PostgresCustomerProvider) FindByPhone(phone string) (*models.Customer, error) {
	database := repo.db
	customer := new(models.Customer)
	result := database.Where("phone_number = ?", phone).First(&customer)
	return customer, result.Error
}

func (repo *PostgresCustomerProvider) FindByID(id string) (*models.Customer, error) {
	database := repo.db
	customer := new(models.Customer)
	result := database.Where("customer_id = ?", id).First(&customer)
	return customer, result.Error
}

func (repo *PostgresCustomerProvider) FilterByOptions(optionsKey string, optionsValue interface{}, per_page, pg string) (customers []models.Customer, page, totalPage int, err error) {
	database := repo.db

	pagination := helpers.CreatePagination(per_page, pg, "")
	customers = make([]models.Customer, 0)

	cursor := database.
		Scopes(helpers.Paginate(&models.Customer{}, pagination, database)).
		Where(optionsKey+" = ?", optionsValue).
		Find(&customers)

	return customers, pagination.GetPage(), pagination.TotalPages, cursor.Error

}
