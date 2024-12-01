package controller_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("AddProduct", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/add-product", ctrl.AddProduct)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":1}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/add-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAddProduct_InvalidPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/add-product", ctrl.AddProduct)

	payload := `{"client_id":"invalid-uuid","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":1}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/add-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAddProduct_Err(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.
		On("AddProduct", mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("error"))
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/add-product", ctrl.AddProduct)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":1}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/add-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestEditProduct_InvalidPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/edit-product", ctrl.EditProduct)

	payload := `{"client_id":"invalid-uuid","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":1}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/edit-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestEditProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("EditProduct", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/edit-product", ctrl.EditProduct)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":2}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/edit-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestEditProduct_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("EditProduct", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/edit-product", ctrl.EditProduct)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":2}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/edit-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestRemoveProduct_InvalidPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/remove-product", ctrl.RemoveProduct)

	payload := `{"client_id":"invalid-uuid","product_id":"123e4567-e89b-12d3-a456-426614174001","quantity":1}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/remove-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestRemoveProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/remove-product", ctrl.RemoveProduct)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000","product_id":"123e4567-e89b-12d3-a456-426614174001"}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/remove-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestRemoveProduct_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/remove-product", ctrl.RemoveProduct)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000","product_id":"123e4567-e89b-12d3-a456-426614174001"}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/remove-product", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestOverview_InvalidPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/overview", ctrl.Overview)

	payload := `{"client_id":"invalid-uuid"}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/overview", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestOverview_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("GetFullCart", mock.Anything).Return(nil, nil)
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/overview", ctrl.Overview)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000"}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/overview", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestOverview_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := new(mocks.IService)
	mockService.On("GetFullCart", mock.Anything).Return(nil, errors.New("error"))
	ctrl := controller.NewCartController(mockService)
	router.POST("/cart/overview", ctrl.Overview)

	payload := `{"client_id":"123e4567-e89b-12d3-a456-426614174000"}`
	req, _ := http.NewRequest(http.MethodPost, "/cart/overview", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}
