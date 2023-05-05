package initdata

import (
	"git.cyradar.com/license-manager/backend/internal/helpers"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"github.com/google/uuid"
)

func InitAccount() {
	salt := helpers.GenerateSalt(16)
	password := helpers.HashPassword(DefaultPassword, salt)

	timeDefault := helpers.ParseDate("2017-06-06")
	id := uuid.NewString()
	pw := string(password)
	saltString := string(salt)
	var defaultUser = &models.User{
		UserID:       &id,
		Email:        &DefaultEmail,
		Password:     &pw,
		Salt:         &saltString,
		Name:         &DefaultName,
		Role:         &DefaultRole,
		Status:       &DefaultStatus,
		LastLoggedIn: timeDefault,
		CreatedAt:    timeDefault,
		UpdatedAt:    timeDefault,
	}
	_, err := persistence.User().SearchByEmail(*defaultUser.Email)
	if err != nil {
		persistence.User().Save(*defaultUser)
	}
}
