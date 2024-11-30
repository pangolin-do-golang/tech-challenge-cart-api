package controller_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchProducts_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IProductService)
	mockService.On("Search", mock.Anything, mock.Anything).Return(nil, nil)
	ctrl := controller.NewProductController(mockService)
	router.GET("/product", ctrl.Search)

	req, _ := http.NewRequest(http.MethodGet, "/product?search=phone&category=electronics", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IProductService)
	mockService.On("Delete", mock.Anything).Return(nil)
	ctrl := controller.NewProductController(mockService)
	router.DELETE("/product/:id", ctrl.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/product/123e4567-e89b-12d3-a456-426614174000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}
}

func TestDeleteProduct_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IProductService)
	mockService.On("Search", mock.Anything).Return(errors.New("error"))
	ctrl := controller.NewProductController(mockService)
	router.DELETE("/product/:id", ctrl.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/product/invalid-uuid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
