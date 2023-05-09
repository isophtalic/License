package service

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	customError "github.com/isophtalic/License/internal/error"
	"github.com/isophtalic/License/internal/models"
	"github.com/isophtalic/License/internal/persistence"
)

func GetLicenses(query url.Values) (licenses []models.License, page int, total_pages int) {
	perPge := strings.TrimSpace(query.Get("per_page"))
	pge := strings.TrimSpace(query.Get("page"))
	sort := strings.TrimSpace(query.Get("sort"))
	search := strings.TrimSpace(query.Get("search"))
	option := strings.TrimSpace(query.Get("option"))
	switch {
	case search != "":
		licenses, page, total_pages = searchByOptions(search, option, perPge, pge, sort)
		break
	case search == "" && option != "":
		licenses, page, total_pages = filterByOptions("option", option, perPge, pge)
		break
	default:
		licenses, page, total_pages = getLicenses(perPge, pge, sort)
		break
	}
	return
}

func GetLicenseByID(id string) (*models.License, error) {
	return persistence.License().FindById(id)
}

func searchByOptions(valueSearch, valueFilter, perPge, pge, sort string) (licenses []models.License, page int, total_pages int) {

	licenses, page, total_pages, err := persistence.License().SearchOrFilter(valueSearch, valueFilter, perPge, pge, sort)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return
	}

	return
}

func filterByOptions(options string, order interface{}, perPge string, pge string) (licenses []models.License, page int, total_pages int) {

	ok := verifyOrderFilter(order.(string))
	if !ok {
		customError.Throw(http.StatusUnprocessableEntity, errors.New("Can't execute order: "+order.(string)).Error())
		return
	}

	licenses, page, total_pages, err := persistence.License().FindByOptions(options, order, perPge, pge)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return
	}

	return
}

func getLicenses(perPge, pge, sort string) (licenses []models.License, page int, total_pages int) {
	licenses, page, total_pages, err := persistence.License().GetLicenses(perPge, pge, sort)
	if err != nil {
		customError.Throw(http.StatusNotFound, "Not Found")
	}
	return
}

func verifyOrderFilter(order string) bool {
	if order == "Pro" || order == "Trial" {
		return true
	}
	return false
}
