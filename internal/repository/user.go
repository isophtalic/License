package repository

import (
	"git.cyradar.com/license-manager/backend/internal/models"
)

type UserRepository interface {
	GetUsers(per_page, pg, sort string) (users []models.User, page, totalPage int, err error)
	Save(user models.User) error

	UpdateProfile(user models.User) error
	UpdateLastLogin(user models.User) error

	FindByID(id string) (*models.User, error)
	SearchOrFilter(valueSearch, valueFilter, per_page, pg, sort string) (users []models.User, page, totalPage int, err error)
	SearchByEmail(email string) (*models.User, error)

	FilterByOptions(optionsKey string, optionsValue interface{}, per_page, pg string) (users []models.User, page, totalPage int, err error)
}
