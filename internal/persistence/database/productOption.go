package database

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ProductOptionDatabase *PostgresPdOptionProvider

type PostgresPdOptionProvider struct {
	table string
	db    *gorm.DB
}

func NewPostgresPdOptionProvider(tableName string, db interface{}) *PostgresPdOptionProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()
	ProductOptionDatabase = &PostgresPdOptionProvider{
		table: tableName,
		db:    database,
	}

	return ProductOptionDatabase
}

func (repo *PostgresPdOptionProvider) GetDB() *gorm.DB {
	return repo.db
}

func (repo *PostgresPdOptionProvider) FindByName(name string) *models.ProductOption {
	var option *models.ProductOption
	err := repo.db.Model(&models.ProductOption{}).Where("name", name).First(&option).Error
	if err != nil {
		return nil
	}
	return option
}

func (repo *PostgresPdOptionProvider) FindByID(id string) *models.ProductOption {
	var option *models.ProductOption
	err := repo.db.Model(&models.ProductOption{}).Where("option_id", id).First(&option).Error
	if err != nil {
		customError.Throw(http.StatusBadRequest, fmt.Sprintf("Error: Find option product by name:::%v", err))
	}
	return option
}

func (repo *PostgresPdOptionProvider) Create(option *dto.ProductOptionDTO) {
	productID := *option.ProductID_FK
	product := ProductDatabase.FindByID(productID)

	if product == nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("Product with ID: %v not found.", productID))
	}

	optionID := uuid.NewString()
	t := time.Now()
	_o := models.ProductOption{
		OptionID:     &optionID,
		Name:         option.Name,
		Description:  option.Description,
		Enable:       new(bool),
		CreatedAt:    &t,
		UpdatedAt:    &t,
		CreatorID_FK: option.CreatorID_FK,
		ProductID_FK: &productID,
	}

	err := repo.db.Transaction(func(tx *gorm.DB) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}
		}()

		if e1 := tx.Create(&_o).Error; e1 != nil {
			println(fmt.Sprintf("Error: Create option:::%v", e1))
			panic(e1)
		}
		ProductOptionDetailDatabase.WithTransaction(tx).Create(optionID)
		return err
	})

	if err != nil {
		customError.Throw(http.StatusBadRequest, fmt.Sprintf("Cannot create options:::%v", err))
	}
}

func (repo *PostgresPdOptionProvider) UpdateByID(id string, updatedData *dto.ProductOptionDTO) {
	option := repo.FindByID(id)
	if option == nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("Option with ID: %v not found.", id))
	}
	//
	// Need to update option detail along with option product ???
	//
	r := repo.db.Where("option_id", id).Updates(dto.ToProductOption(updatedData))
	if r.Error != nil {
		println(fmt.Sprintf("Error: Update product option by ID:::%v", r.Error))
		panic(r.Error)
	}
}

func (repo *PostgresPdOptionProvider) DeleteByID(id string) *models.ProductOption {
	var deletedOption models.ProductOption
	r := repo.db.Clauses(clause.Returning{
		Columns: []clause.Column{{Name: "name"}, {Name: "description"}},
	}).Where("option_id", id).Delete(&deletedOption)
	if r.Error != nil {
		println(fmt.Sprintf("Error: Delete product option by ID:::%v", r.Error))
		panic(r.Error)
	}
	return &deletedOption
}
