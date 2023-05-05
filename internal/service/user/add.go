package service

import (
	"fmt"
	"time"

	"git.cyradar.com/license-manager/backend/internal/dto"
	"git.cyradar.com/license-manager/backend/internal/helpers"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func AddUser(token string, cmd *dto.AddUserDTO) error {

	_, role, err := helpers.DecodeToken(token)
	if err != nil {
		return err
	}
	if role != "admin" {
		return errors.New("Permission denied")
	}

	validators.ValidateUser(cmd)

	_, err = persistence.User().SearchByEmail(cmd.Email)
	if err != nil && err.Error() != "record not found" {
		return errors.New(fmt.Sprintf("User %s existed", cmd.Email))
		// return err
	}
	salt := helpers.GenerateSalt(16)
	password := helpers.HashPassword(cmd.Password, salt)
	pw := string(password)
	id := uuid.NewString()
	saltString := string(salt)
	uUser := models.User{
		UserID:       &id,
		Email:        &cmd.Email,
		Password:     &pw,
		Salt:         &saltString,
		Name:         &cmd.Name,
		Role:         &cmd.Role,
		Status:       &cmd.Status,
		LastLoggedIn: time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return persistence.User().Save(uUser)
}
