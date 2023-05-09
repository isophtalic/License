package service

import (
	"errors"
	"time"

	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/validators"
)

func UpdateProfile(cmd *dto.UpdateUserDTO, email, id string) error {
	var user *models.User
	var err error
	validators.ValidateUpdateProfile(cmd)

	if id != "" {
		user, err = persistence.User().FindByID(id)
		if err != nil {
			return err
		}
		if *user.Role == "admin" {
			return errors.New("Can't update for another admin")
		}
	} else {
		user, err = persistence.User().SearchByEmail(email)
		if err != nil {
			return err
		}
		if cmd.Email != email {
			_, err = persistence.User().SearchByEmail(cmd.Email)
			if err == nil {
				return errors.New("Email already used")
			}
		}
	}

	userUpdate := models.User{
		UserID:       user.UserID,
		Email:        &cmd.Email,
		Password:     user.Password,
		Salt:         user.Salt,
		Name:         &cmd.Name,
		Role:         &cmd.Role,
		Status:       &cmd.Status,
		LastLoggedIn: user.LastLoggedIn,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	return persistence.User().UpdateProfile(userUpdate)
}

func ChangePassword(data *dto.ChangePassword, id, token string) error {
	var email, role string
	var user *models.User
	err := validators.ValidateFormChangePassword(data)
	if err != nil {
		return err
	}

	email, role, err = helpers.DecodeToken(token)
	if err != nil {
		return err
	}
	if id != "" {
		if role != "admin" {
			return errors.New("Permission denied")
		}
		user, err = persistence.User().FindByID(id)
		if err != nil {
			return err
		}
		if *user.Role == "admin" && *user.Email != email {
			return errors.New("Can't change password for another admin")
		}
	} else {
		user, err = persistence.User().SearchByEmail(email)
		if err != nil {
			return err
		}
		//
	}

	if err := validators.VerifyPassword(data.Password, data.NewPassword, *user.Salt, *user.Password); err != nil {
		return err
	}

	password := helpers.HashPassword(data.NewPassword, *user.Salt)

	if data.ConfirmPassword != data.NewPassword {
		return errors.New("Require enter confirm_password equal to new_password")
	}

	if role == "admin" {
		return handleChangePassword(user, string(password))
	}

	return handleChangePassword(user, string(password))
}

func handleChangePassword(user *models.User, hashPassword string) error {
	userUpdate := models.User{
		UserID:       user.UserID,
		Email:        user.Email,
		Password:     &hashPassword,
		Salt:         user.Salt,
		Name:         user.Name,
		Role:         user.Role,
		Status:       user.Status,
		LastLoggedIn: user.LastLoggedIn,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    time.Now(),
	}
	err := LogOut(*user.UserID + "*")
	if err != nil {
		return errors.New("Cannot delete token in database")
	}
	return persistence.User().UpdateProfile(userUpdate)
}
