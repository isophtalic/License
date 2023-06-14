package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/helpers"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
	"gorm.io/gorm"
)

func findProducts(kw string, query url.Values, attrs []string) ([]models.Product, int, int) {
	perPage := strings.TrimSpace(query.Get("per_page"))
	page := strings.TrimSpace(query.Get("page"))
	sort := strings.TrimSpace(query.Get("sort"))
	status := strings.TrimSpace(query.Get("status"))
	pagination := helpers.CreatePagination(perPage, page, sort)

	q := persistence.Product().GetDB().Preload("Creator", func(db *gorm.DB) *gorm.DB {
		return db.Select("name", "user_id")
	}).Scopes(helpers.Paginate(&models.Product{}, pagination, persistence.Product().GetDB()))

	if status != "" {
		s, err := strconv.ParseBool(status)
		if err != nil {
			customError.Throw(http.StatusBadRequest, fmt.Sprintf("Invalid filter status:::%v", err))
		}
		q = q.Where(&models.Product{Status: &s})
	}

	if strings.TrimSpace(kw) != "" {
		q = q.Where("email LIKE ? OR company LIKE ? OR phone LIKE ? OR address LIKE ? OR description LIKE ?", kw, kw, kw, kw, kw)
	}

	var products []models.Product
	r := q.
		Select(attrs).
		Find(&products)

	if r.Error != nil {
		customError.Throw(http.StatusBadRequest, r.Error.Error())
	}
	return products, pagination.GetPage(), pagination.GetTotalPages()
}

func ListProducts(query url.Values) ([]models.Product, int, int) {
	return findProducts("", query, []string{"product_id", "name", "status", "created_at", "creator_id"})
}

func DetailProduct(id string) *models.Product {
	var product models.Product
	err := persistence.Product().GetDB().
		Preload("Key", func(db *gorm.DB) *gorm.DB {
			return db.Select("key_id", "type", "public_key", "product_id", "creator_id")
		}).
		Preload("Key.Creator", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id", "name")
		}).
		Where("product_id", id).Omit("updated_at", "creator_id").First(&product).Error

	if err != nil {
		println(fmt.Println("Detail pd:::", err))
		customError.Throw(http.StatusBadRequest, err.Error())
	}
	return &product
}

func GetProductOptions(productID string, query url.Values) ([]models.ProductOption, int, int) {
	perPage := strings.TrimSpace(query.Get("per_page"))
	page := strings.TrimSpace(query.Get("page"))
	sort := strings.TrimSpace(query.Get("sort"))
	pagination := helpers.CreatePagination(perPage, page, sort)

	var options []models.ProductOption

	r := persistence.ProductOption().GetDB().
		Preload("Creator", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id", "name")
		}).
		Scopes(helpers.Paginate(&models.ProductOption{}, pagination, persistence.ProductOption().GetDB())).
		Where("product_id", productID).
		Select("option_id", "name", "enable", "created_at", "creator_id").
		Find(&options)

	println(fmt.Printf("%v", options))

	if r.Error != nil {
		println(fmt.Printf("Detail pd:::%v", r.Error))
		customError.Throw(http.StatusBadRequest, r.Error.Error())
	}
	return options, pagination.GetPage(), pagination.GetTotalPages()
}

func Search(keyWord string, query url.Values) ([]models.Product, int, int) {
	if strings.TrimSpace(keyWord) == "" {
		return nil, 0, 0
	}
	keyword := "%" + strings.TrimSpace(keyWord) + "%"
	return findProducts(keyword, query, []string{"product_id", "name", "status", "email", "description", "created_at", "creator_id"})
}

func GetKey(productID string) (keyRes *keyResponse) {
	keyDetail, err := persistence.Key().GetKeyByProductID(productID)
	if err != nil {
		customError.Throw(http.StatusUnprocessableEntity, "Don't get product key by :"+err.Error())
	}
	keyRes = &keyResponse{
		Type:        keyDetail.Type,
		CreatedAt:   keyDetail.CreatedAt,
		CreatorName: keyDetail.Creator.Name,
	}
	return keyRes
}

type keyResponse struct {
	Type        *string
	CreatedAt   *time.Time
	CreatorName *string
}
