package handler

import (
	"subscribe-test-service/docs"
	"subscribe-test-service/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	subscriptions := router.Group("/subscriptions")
	{
		subscriptions.POST("/", h.createSubscription)
		subscriptions.GET("/", h.getAllSubscriptions)
		subscriptions.GET("/:id", h.getSubscription)
		subscriptions.PUT("/:id", h.updateSubscription)
		subscriptions.DELETE("/:id", h.deleteSubscription)
	}

	return router
}
