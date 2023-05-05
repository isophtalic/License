package service

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/helpers"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
)

func GetUsers(token string, query url.Values) (users []models.User, page int, total_pages int) {
	perPge := strings.TrimSpace(query.Get("per_page"))
	pge := strings.TrimSpace(query.Get("page"))
	sort := strings.TrimSpace(query.Get("sort"))
	search := strings.TrimSpace(query.Get("search"))
	status := strings.TrimSpace(query.Get("status"))
	permission := strings.TrimSpace(query.Get("role"))

	_, role, err := helpers.DecodeToken(token)
	if err != nil {
		customError.Throw(http.StatusBadRequest, err.Error())
	}
	if role != "admin" {
		customError.Throw(http.StatusMethodNotAllowed, errors.New("Permission denied").Error())
	}

	switch {
	case search != "" && status == "" && permission == "":
		users, page, total_pages = searchByOptions(search, "", perPge, pge, sort)
		break
	case search != "" && status != "" && permission == "":
		users, page, total_pages = searchByOptions(search, status, perPge, pge, sort)
		break
	case search != "" && status == "" && permission != "":
		users, page, total_pages = searchByOptions(search, permission, perPge, pge, sort)
		break
	case search == "" && status != "" && permission == "":
		users, page, total_pages = filterByOptions("status", status, perPge, pge)
		break
	case search == "" && status == "" && permission != "":
		users, page, total_pages = filterByOptions("role", permission, perPge, pge)
		break
	default:
		users, page, total_pages = getUsers(perPge, pge, sort)
		break
	}
	return
}

func Profile(email string) *models.User {
	user, err := persistence.User().SearchByEmail(email)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return &models.User{}
	}
	return user
}

func filterByOptions(options string, order interface{}, perPge string, pge string) (users []models.User, page int, total_pages int) {

	if order != "true" && order != "false" && order != "admin" && order != "customer" {
		customError.Throw(http.StatusUnprocessableEntity, errors.New("Can't execute order: "+order.(string)).Error())
		return
	}

	users, page, total_pages, err := persistence.User().FilterByOptions(options, order, perPge, pge)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return
	}

	return
}

func searchByOptions(valueSearch, valueFilter, perPge, pge, sort string) (users []models.User, page int, total_pages int) {

	users, page, total_pages, err := persistence.User().SearchOrFilter(valueSearch, valueFilter, perPge, pge, sort)
	if err != nil {
		customError.Throw(http.StatusNotFound, err.Error())
		return
	}

	return
}

func getUsers(perPge, pge, sort string) (users []models.User, page int, total_pages int) {
	users, page, total_pages, err := persistence.User().GetUsers(perPge, pge, sort)
	if err != nil {
		customError.Throw(http.StatusNotFound, errors.New("Not Found").Error())
	}

	return
}
