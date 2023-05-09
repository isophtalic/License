package postgres

import (
	"fmt"

	"github.com/isophtalic/License/internal/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgresQL(config *configs.Configure) *Postgres {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	fmt.Println("Connecting DB . . .")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("Connect postgres fail " + err.Error())
	}

	return &Postgres{
		db: db,
	}
}

func (pg *Postgres) GetDB() *gorm.DB {
	return pg.db
}
