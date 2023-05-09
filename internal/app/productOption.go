package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	productOptionService "github.com/isophtalic/License/internal/service/productOption"
)

// [GET] /product-option/:id
func DetailProductOption() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")
		var productOption dto.ProductOptionDTO
		err := c.ShouldBindJSON(&productOption)
		if err != nil {
			customError.Throw(http.StatusBadRequest, err.Error())
		}
		optionDetail := productOptionService.DetailProductOption(id)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Add options successfully",
			Data:    optionDetail,
		})
	})
}

// [POST] /product-option
func CreateOptions() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		creatorEmail := c.Request.Context().Value("user.email").(string)
		var productOption dto.ProductOptionDTO
		err := c.ShouldBindJSON(&productOption)
		if err != nil {
			customError.Throw(http.StatusBadRequest, err.Error())
		}
		productOptionService.CreateOptions(creatorEmail, &productOption)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Add options successfully",
		})
	})
}

// [PATCH] /product-option/:id
func UpdateProductOption() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		optionID := c.Param("id")
		var optionBody dto.ProductOptionDTO
		err := c.ShouldBindJSON(&optionBody)
		if err != nil {
			customError.Throw(http.StatusBadRequest, "Invalid body data")
		}
		productOptionService.UpdateProductOption(optionID, &optionBody)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Update option successfully",
		})
	})
}

// [DELETE] /product-option/:id
func DeleteProductOption() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")
		option := productOptionService.DeleteOptionByID(id)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Add option detail successfully",
			Data:    option,
		})
	})
}

// [POST] /product-option/:id/detail/create
func AddOptionDetail() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		optionID := c.Param("id")
		var optionDetail dto.OptionDetailDTO
		err := c.ShouldBindJSON(&optionDetail)
		if err != nil {
			customError.Throw(http.StatusBadRequest, "Invalid body data")
		}
		option := productOptionService.AddOptionDetail(optionID, optionDetail)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Add option detail successfully",
			Data:    option,
		})
	})
}
