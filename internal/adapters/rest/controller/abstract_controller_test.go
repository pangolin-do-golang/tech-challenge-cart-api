package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/errutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAbstractController_Error_BusinessError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	ctrl := &AbstractController{}
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, &errutil.Error{Type: "BUSINESS", Message: "Business error"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, got %d", w.Code)
	}
	if w.Body.String() != `{"error":"Business error"}` {
		t.Errorf("Expected body Business error, got %s", w.Body.String())
	}
}

func TestAbstractController_Error_InputError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	ctrl := &AbstractController{}
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, &errutil.Error{Type: "INPUT"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
	if w.Body.String() != `{"error":"Bad Request"}` {
		t.Errorf("Expected body Bad Request, got %s", w.Body.String())
	}
}

func TestAbstractController_Error_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	ctrl := &AbstractController{}
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, &errutil.Error{Type: "UNKNOWN"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
	if w.Body.String() != `{"error":"Internal Server Error"}` {
		t.Errorf("Expected body Internal Server Error, got %s", w.Body.String())
	}
}

func TestAbstractController_Error_GenericError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	ctrl := &AbstractController{}
	router.GET("/error", func(c *gin.Context) {
		ctrl.Error(c, errors.New("generic error"))
	})

	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
	if w.Body.String() != `{"error":"Internal Server Error"}` {
		t.Errorf("Expected body Internal Server Error, got %s", w.Body.String())
	}
}
