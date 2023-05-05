package repository

import (
	"git.cyradar.com/license-manager/backend/internal/models"
	"gorm.io/gorm"
)

type KeyRepository interface {
	GetDB() *gorm.DB
	Create(keys ...*models.Key)
}
