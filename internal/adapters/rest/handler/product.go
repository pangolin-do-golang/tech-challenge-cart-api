package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/product"
)

func RegisterProductHandlers(router *gin.Engine, service *product.Service) {
	productController := controller.NewProductController(service)

	router.GET("/product", productController.Search)

	router.DELETE("/product/:id")
}
