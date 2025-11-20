package handler

import (
	docs "subscribe-test-service/docs"
	"subscribe-test-service/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler strucsubscribet {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//api := router.Group("/api")
	//{
	//
	//
	//}

	return router
}
