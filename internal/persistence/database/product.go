package database

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	postgresDB "github.com/isophtalic/License/internal/persistence/postgres"
	"gorm.io/gorm"
)

var ProductDatabase *PostgresProductProvider

type PostgresProductProvider struct {
	table string
	db    *gorm.DB
}

func NewPostgresProductProvider(tableName string, db interface{}) *PostgresProductProvider {
	database := (db.(*postgresDB.Postgres)).GetDB()

	ProductDatabase = &PostgresProductProvider{
		table: tableName,
		db:    database,
	}

	return ProductDatabase
}

func (repo *PostgresProductProvider) GetDB() *gorm.DB {
	return repo.db
}

func (repo *PostgresProductProvider) FindOneByName(name string) *models.Product {
	var product models.Product
	r := repo.db.Model(&models.Product{}).Where(&models.Product{Name: &name}).First(&product)
	if r.Error != nil {
		return nil
	}
	return &product
}

func (repo *PostgresProductProvider) FindByID(id string) *models.Product {
	var product models.Product
	r := repo.db.Model(&models.Product{}).Where(&models.Product{ProductID: &id}).First(&product)
	if r.Error != nil {
		println(fmt.Sprintf("Error:: Find one by id product: %v", r.Error))
		customError.Throw(http.StatusBadRequest, r.Error.Error())
	}
	return &product
}

func (repo *PostgresProductProvider) FindAll(pagination *dto.PaginationDTO) (products []models.Product) {
	r := repo.db.Model(&models.Product{}).
		Scopes(helpers.Paginate(models.Product{}, pagination, repo.db)).
		Find(&products)
	if r.Error != nil {
		customError.Throw(http.StatusBadRequest, fmt.Sprintf("Error:: Find all product: %v", r.Error.Error()))
	}
	return products
}

func (repo *PostgresProductProvider) CreateOne(creatorEmail string, productDTO *dto.ProductDTO) string {
	creator, _err := UserDatabase.SearchByEmail(creatorEmail)
	if _err != nil {
		customError.Throw(http.StatusNotFound, fmt.Sprintf("Can not find user with email: %v", creatorEmail))
	}

	product := models.Product{
		ProductID:    &[]string{uuid.NewString()}[0],
		Name:         productDTO.Name,
		Description:  productDTO.Description,
		Status:       new(bool),
		Company:      productDTO.Company,
		Email:        productDTO.Email,
		Phone:        productDTO.Phone,
		Address:      productDTO.Address,
		CreatedAt:    &[]time.Time{time.Now()}[0],
		UpdatedAt:    &[]time.Time{time.Now()}[0],
		CreatorID_FK: creator.UserID,
	}
	r := repo.db.Create(&product)
	if r.Error != nil {
		println(fmt.Printf("%v", product))
		customError.Throw(http.StatusInternalServerError, fmt.Sprintf("Error:: Create product: %v", r.Error.Error()))
	}
	return *product.ProductID
}

func (repo *PostgresProductProvider) Update(id string, productUpdate *dto.ProductDTO) *models.Product {
	r := repo.db.Where("product_id", id).Updates(dto.ToProduct(productUpdate))
	if r.Error != nil {
		println(fmt.Printf("Error: Update product::: %v", r.Error))
		return nil
	}
	var updatedProduct *models.Product
	repo.db.First(updatedProduct)
	return updatedProduct
}
