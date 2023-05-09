package routes

import (
	"github.com/isophtalic/License/internal/configs"

	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/middleware"
)

func NewAPIv1(config *configs.Configure, mode string) *gin.Engine {
	gin.SetMode(mode)
	serve := gin.Default()
	serve.Use(middleware.CORS(config))
	v1 := serve.RouterGroup.Group("/api/v1")
	authRouter(v1)

	v1.Use(
		middleware.AuthMiddleware(config.JWT_KEY),
	)

	userRouter(v1)
	profileRouter(v1)
	productRouter(v1)
	productOptionRouter(v1)
	optionDetailRouter(v1)
	customerRouter(v1)
	licenseRouter(v1)

	return serve
}
