package repository

import "github.com/isophtalic/License/internal/models"

type CustomerRepository interface {
	GetCustomers(per_page, pg, sort string) (customer []models.Customer, page, totalPage int, err error)
	Save(customer models.Customer) error

	UpdateCustomer(customer models.Customer) error

	SearchOrFilter(valueSearch string, valueFilter, per_page, pg, sort string) (customers []models.Customer, page, totalPage int, err error)
	FindByID(id string) (*models.Customer, error)
	FindByEmail(email string) (*models.Customer, error)
	FindByPhone(phone string) (*models.Customer, error)
	FindByName(name string) (*models.Customer, error)

	FilterByOptions(optionsKey string, optionsValue interface{}, per_page, pg string) (customer []models.Customer, page, totalPage int, err error)
}
