package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/lib/pq"
)

func UpdateCustomer(customerID string, cmd *dto.UpdateAndAddCustomerDTO) error {

	cus, err := persistence.Customer().FindByID(customerID)
	if err != nil {
		return errors.New("Customer is invalid")
	}

	if *cus.Email != cmd.Email {
		_, err := persistence.Customer().FindByEmail(cmd.Email)
		if err == nil {
			return errors.New("Email was used")
		}
	}

	if *cus.Phone_Number != cmd.PhoneNumber {
		_, err := persistence.Customer().FindByPhone(cmd.PhoneNumber)
		fmt.Println(err)
		if err == nil {
			return errors.New("Phone number was used")
		}
	}

	if *cus.Name != cmd.Name {
		_, err := persistence.Customer().FindByName(cmd.Name)
		if err == nil {
			return errors.New("Name was used")
		}
	}

	birthday, err := time.Parse("2006-01-02", cmd.BirthDay)
	if err != nil {
		return errors.New("Birthday field is wrong format")
	}

	tags, _ := pq.Array(cmd.Tags).Value()

	newCustomer := models.Customer{
		CustomerID:   &customerID,
		Email:        &cmd.Email,
		Name:         &cmd.Name,
		Company:      &cmd.Company,
		Phone_Number: &cmd.PhoneNumber,
		Organization: &cmd.Organization,
		BirthDay:     birthday,
		Address:      &cmd.Address,
		Description:  &cmd.Description,
		Tags:         strings.Fields(tags.(string)),
		Enable:       &cmd.Enable,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return persistence.Customer().UpdateCustomer(newCustomer)
}
