package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"github.com/isophtalic/License/internal/validators"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ProductOptionDetailDatabase *PostgresPdOptionDetailProvider

type PostgresPdOptionDetailProvider struct {
	table string
	db    *gorm.DB
}

func NewPostgresPdOptionDetailProvider(tableName string, db interface{}) *PostgresPdOptionDetailProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()

	ProductOptionDetailDatabase = &PostgresPdOptionDetailProvider{
		table: tableName,
		db:    database,
	}

	return ProductOptionDetailDatabase
}

func (repo *PostgresPdOptionDetailProvider) GetDB() *gorm.DB {
	return repo.db
}

func (repo *PostgresPdOptionDetailProvider) WithTransaction(tx *gorm.DB) *PostgresPdOptionDetailProvider {
	cloneRepo := &PostgresPdOptionDetailProvider{
		table: repo.table,
		db:    tx,
	}
	return cloneRepo
}

func (repo *PostgresPdOptionDetailProvider) FindByKey(key string) *models.OptionDetail {
	var optionDetail models.OptionDetail
	r := repo.db.Where("key", key).First(&optionDetail)
	if r.Error != nil {
		println(fmt.Sprintf("Error: Find option detail by key:::%v", r.Error))
		panic(r.Error)
	}
	return &optionDetail
}

func (repo *PostgresPdOptionDetailProvider) FindByID(id string) *models.OptionDetail {
	var optionDetail models.OptionDetail
	r := repo.db.Where("option_detail_id", id).First(&optionDetail)
	if r.Error != nil {
		println(fmt.Sprintf("Error: Find option detail by id:::%v", r.Error))
		return nil
	}
	return &optionDetail
}

func (repo *PostgresPdOptionDetailProvider) Create(productOptionID string, optionDetails ...dto.OptionDetailDTO) {
	if len(optionDetails) == 0 {
		return
	}

	optionDts := []models.OptionDetail{}
	for _, option := range optionDetails {
		validators.ValidateProductOptionDetail(&option)
		t := time.Now()
		optionDts = append(optionDts, models.OptionDetail{
			OptionDetailID:     &[]string{uuid.NewString()}[0],
			Key:                option.Key,
			Value:              option.Value,
			CreatedAt:          &t,
			UpdatedAt:          &t,
			ProductOptionID_FK: &productOptionID,
		})
	}

	err := repo.db.Create(&optionDts).Error
	if err != nil {
		println(fmt.Printf("Error:Create option detail:::%v", err))
		panic(err)
	}

}

func (repo *PostgresPdOptionDetailProvider) DeleteByID(id string) *models.OptionDetail {
	var deletedOptionDetail models.OptionDetail
	r := repo.db.Clauses(clause.Returning{}).Where("option_detail_id", id).Delete(&deletedOptionDetail)
	if r.Error != nil {
		println(fmt.Sprintf("Error: Delete product option by ID:::%v", r.Error))
		panic(r.Error)
	}
	return &deletedOptionDetail
}
