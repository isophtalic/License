package service

import (
	"fmt"
	"net/http"

	"git.cyradar.com/license-manager/backend/internal/dto"
	customError "git.cyradar.com/license-manager/backend/internal/error"
	"git.cyradar.com/license-manager/backend/internal/models"
	"git.cyradar.com/license-manager/backend/internal/persistence"
	"git.cyradar.com/license-manager/backend/internal/validators"
)

func CreateOptions(creatorEmail string, option *dto.ProductOptionDTO) {
	creator, err := persistence.User().SearchByEmail(creatorEmail)
	if creator == nil || err != nil {
		customError.Throw(http.StatusNotFound, "Something went wrong while finding user.")
	}
	option.CreatorID_FK = creator.UserID
	existingProductOption := persistence.ProductOption().FindByName(*option.Name)
	if existingProductOption != nil {
		customError.Throw(http.StatusConflict, fmt.Sprintf("Option with name: '%v' already existed.ID: '%v'", *option.Name, *existingProductOption.OptionID))
	}
	validators.ValidateProductOption(option)
	persistence.ProductOption().Create(option)
}

func AddOptionDetail(optionID string, optionDetail dto.OptionDetailDTO) *models.OptionDetail {
	// validators.ValidateCreateProductOptionDetail(&optionDetail)
	// option := persistence.ProductOption()..
	// 	Association("OptionDetails", []string{"key", "product_option_id"}).
	// 	FindOneByAttributes(map[string]interface{}{}).
	// 	Value()
	// if option == nil {
	// 	customError.Throw(http.StatusNotFound, "Can not specify option.")
	// }
	// for _, opD := range option.OptionDetails {
	// 	opD.Key = strings.TrimSpace(opD.Key)
	// 	if opD.Key == optionDetail.Key {
	// 		customError.Throw(http.StatusConflict, fmt.Sprintf("Key: '%v' already exist", opD.Key))
	// 	}
	// }
	// optionDetail.ProductOptionID_FK = optionID
	// return persistence.ProductOptionDetail().Create(&optionDetail)
	return nil
}
