package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/cart"
)

func RegisterCartHandlers(router *gin.Engine, service cart.IService) {
	cartController := controller.NewCartController(service)

	router.POST("/cart/overview", cartController.Overview)
	router.POST("/cart/add-product", cartController.AddProduct)
	router.POST("/cart/remove-product", cartController.RemoveProduct)
	router.POST("/cart/edit-product", cartController.EditProduct)
}
