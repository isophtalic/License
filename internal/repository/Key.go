package repository

import (
	"github.com/isophtalic/License/internal/models"
	"gorm.io/gorm"
)

type KeyRepository interface {
	GetDB() *gorm.DB
	Create(keys ...*models.Key)
}
