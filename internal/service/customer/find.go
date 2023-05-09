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

func GetCustomers(query url.Values) (customers []models.Customer, page, total_pages int) {
	perPge := strings.TrimSpace(query.Get("per_page"))
	pge := strings.TrimSpace(query.Get("page"))
	sort := strings.TrimSpace(query.Get("sort"))
	search := strings.TrimSpace(query.Get("search"))
	status := strings.TrimSpace(query.Get("enable"))
	organization := strings.TrimSpace(query.Get("organization"))
	switch {
	case search != "" && status == "" && organization == "":
		customers, page, total_pages = searchByOptions(search, "", perPge, pge, sort)
		break
	case search != "" && status != "" && organization == "":
		customers, page, total_pages = searchByOptions(search, status, perPge, pge, sort)
		break
	case search != "" && status == "" && organization != "":
		customers, page, total_pages = searchByOptions(search, organization, perPge, pge, sort)
		break
	case search == "" && status != "" && organization == "":
		customers, page, total_pages = filterByOptions("status", status, perPge, pge)
		break
	case search == "" && status == "" && organization != "":
		customers, page, total_pages = filterByOptions("organization", organization, perPge, pge)
		break
	default:
		customers, page, total_pages = getCustomers(perPge, pge, sort)
		break
	}
	return
}

func filterByOptions(options string, order interface{}, perPge string, pge string) (customers []models.Customer, page int, total_pages int) {

	ok := verifyOrderFilter(order.(string))
	if !ok {
		customError.Throw(http.StatusUnprocessableEntity, errors.New("Can't execute order: "+order.(string)).Error())
		return
	}

	customers, page, total_pages, err := persistence.Customer().FilterByOptions(options, order, perPge, pge)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return
	}

	return
}

func searchByOptions(valueSearch, valueFilter, perPge, pge, sort string) (customer []models.Customer, page int, total_pages int) {

	customer, page, total_pages, err := persistence.Customer().SearchOrFilter(valueSearch, valueFilter, perPge, pge, sort)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return
	}

	return
}

func getCustomers(perPge, pge, sort string) (customers []models.Customer, page int, total_pages int) {
	customers, page, total_pages, err := persistence.Customer().GetCustomers(perPge, pge, sort)
	if err != nil {
		customError.Throw(http.StatusNotFound, "Not Found")
	}
	return
}

func verifyOrderFilter(order string) bool {
	if order == "true" || order == "false" || order == "personal" || order == "government" || order == "enterprise" {
		return true
	}
	return false
}
