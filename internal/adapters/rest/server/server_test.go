package server_test

import (
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/rest/server"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint_ReturnsOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestServe_StartsServer(t *testing.T) {
	productService := new(mocks.IProductService)
	cartService := new(mocks.IService)
	rs := server.NewRestServer(&server.RestServerOptions{
		ProductService: productService,
		CartService:    cartService,
	})

	go func() {
		rs.Serve()
	}()
}

func TestRegisterProductHandlers_CallsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	productService := new(mocks.IProductService)
	productService.On("Search", mock.Anything, mock.Anything).Return(nil, nil)
	handler.RegisterProductHandlers(router, productService)

	req, _ := http.NewRequest(http.MethodGet, "/product", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterCartHandlers_CallsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	cartService := new(mocks.IService)
	cartService.On("GetFullCart").Return(nil, nil)
	handler.RegisterCartHandlers(router, cartService)

	req, _ := http.NewRequest(http.MethodPost, "/cart/overview", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
