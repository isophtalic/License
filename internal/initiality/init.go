package initdata

import (
	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
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
