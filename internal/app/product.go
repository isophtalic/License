package app

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	productService "github.com/isophtalic/License/internal/service/product"
)

// [POST] /product
func CreateProduct() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		creatorEmail := c.Request.Context().Value("user.email").(string)
		var data dto.ProductDTO
		err := c.ShouldBindJSON(&data)
		if err != nil {
			customError.Throw(http.StatusBadRequest, "Something went wrong while reading body.")
		}
		id := productService.Create(creatorEmail, &data)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Create product successfully",
			Data:    id,
		})
	})
}

// [GET] /product?per_page&page&status&sort
func ListProducts() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		query := c.Request.URL.Query()
		products, page, totalPages := productService.ListProducts(query)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Get product successfully",
			Data: map[string]interface{}{
				"page":       page,
				"totalPages": totalPages,
				"products":   products,
			},
		})
	})
}

// [PATCH] /product/:id
func UpdateProduct() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")
		var data dto.ProductDTO
		err := c.ShouldBindJSON(&data)
		if err != nil {
			Response(c, http.StatusBadRequest, ResponseBody{
				Message: err.Error(),
			})
		}
		updatedProduct := productService.Update(id, &data)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Update product successfully",
			Data:    updatedProduct,
		})
	})
}

// [PATCH] /product/status/:id
func ChangeProductStatus() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")
		var data struct {
			Status bool `json:"status"`
		}
		err := c.ShouldBindJSON(&data)
		if err != nil {
			customError.Throw(http.StatusBadRequest, err.Error())
		}
		productService.ChangeStatus(id, data.Status)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Update status product successfully",
		})
	})
}

// [GET] /product/:id
func DetailProduct() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		id := c.Param("id")
		product := productService.DetailProduct(id)
		Response(c, http.StatusOK, ResponseBody{
			Data: product,
		})
	})
}

// [GET] /product/search/:keyword&status&sort
func Search() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		keyWord := c.Param("keyword")
		query := c.Request.URL.Query()
		products, page, totalPages := productService.Search(keyWord, query)
		Response(c, http.StatusOK, ResponseBody{
			Data: map[string]interface{}{
				"options":    products,
				"page":       page,
				"totalPages": totalPages,
			},
		})
	})
}

// [GET] /product/:id/key?type
func GenerateNewKeys() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		productID := c.Param("id")
		typeKey := c.Request.URL.Query().Get("type")
		creatorEmail := c.Request.Context().Value("user.email").(string)
		productService.GenerateKeys(productID, creatorEmail, typeKey)
		Response(c, http.StatusOK, ResponseBody{
			Message: "Generate keys successfully",
		})
	})
}

// [POST] /product/:id/key/upload?type
func UploadKeys() gin.HandlerFunc {
	return HandleErrorWrapper(func(c *gin.Context) {
		productID := c.Param("id")
		typeKey := c.Request.URL.Query().Get("type")
		creatorEmail := c.Request.Context().Value("user.email").(string)
		fileKeys, exist := c.Get("file_keys")
		if !exist {
			customError.Throw(http.StatusBadRequest, "Some thing went wrong while reading body.")
		}
		productService.UploadKeys(productID, creatorEmail, typeKey, fileKeys.([]*multipart.FileHeader))
		Response(c, http.StatusOK, ResponseBody{
			Message: "Upload keys successfully",
		})
	})
}
