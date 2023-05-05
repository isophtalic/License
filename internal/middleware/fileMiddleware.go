package middleware

import (
	"fmt"
	"net/http"
	"path/filepath"

	customError "git.cyradar.com/license-manager/backend/internal/error"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

/*
	request more than size(MB) will be aborted
*/
func SizeLimiterMiddleware(size int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > size {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"message": fmt.Sprintf("Data too large(%v MB). Max: %v MB", c.Request.ContentLength, size),
			})
			c.Abort()
		}
		c.Next()
	}
}

func ExtRestrictMiddleware(ext []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			customError.Throw(http.StatusBadRequest, err.Error())
		}
		files := form.File["keys"]
		for _, file := range files {
			fileExt := filepath.Ext(file.Filename)
			if !slices.Contains(ext, fileExt) {
				c.JSON(http.StatusNotImplemented, gin.H{
					"message": fmt.Sprintf("File type: '%v' is not allowed. Only: %v", fileExt, ext),
				})
				c.Abort()
			}
		}
		c.Set("file_keys", files)
		c.Next()
	}
}
