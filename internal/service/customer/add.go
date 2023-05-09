package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func AddCustomer(cmd *dto.UpdateAndAddCustomerDTO) error {
	validators.ValidateCustomer(cmd)
	_, err := persistence.Customer().FindByEmail(cmd.Email)
	if err == nil {
		return errors.New("Email was used")
	}

	_, err = persistence.Customer().FindByPhone(cmd.PhoneNumber)
	if err == nil {
		return errors.New("Phone number was used")
	}
	_, err = persistence.Customer().FindByName(cmd.Name)
	if err == nil {
		return errors.New("Name was used")
	}

	id := uuid.NewString()
	birthday, err := time.Parse("2006-01-02", cmd.BirthDay)
	if err != nil {
		return errors.New("Birthday field is wrong format")
	}

	// tags, _ := pq.Array(cmd.Tags).Value()

	newCustomer := models.Customer{
		CustomerID:   &id,
		Email:        &cmd.Email,
		Name:         &cmd.Name,
		Company:      &cmd.Company,
		Phone_Number: &cmd.PhoneNumber,
		Organization: &cmd.Organization,
		BirthDay:     birthday,
		Address:      &cmd.Address,
		Description:  &cmd.Description,
		Tags:         cmd.Tags,
		Enable:       &cmd.Enable,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return persistence.Customer().Save(newCustomer)
}

// type StringArray []string

// func (a StringArray) Value() (driver.Value, error) {
// 	return strings.Join(a, ","), nil
// }
