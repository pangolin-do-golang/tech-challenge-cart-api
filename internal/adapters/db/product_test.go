package db_test

import (
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/product"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetByID_ReturnsProduct(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresProductRepository(mockDB)

	productID := uuid.New()
	expectedProduct := &product.Product{
		Id:          productID,
		Name:        "Test Product",
		Description: "Test Description",
		Category:    "Test Category",
		Price:       100.0,
	}

	mockDB.On("First", mock.Anything, "id = ? and deleted_at is null", productID).Return(noErr).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*db.ProductPostgres)
		*arg = db.ProductPostgres{
			BaseModel:   db.BaseModel{ID: productID},
			Name:        "Test Product",
			Description: "Test Description",
			Category:    "Test Category",
			Price:       100.0,
		}
	})

	result, err := repo.GetByID(productID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, result)
	mockDB.AssertExpectations(t)
}

func TestGetByID_ReturnsError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresProductRepository(mockDB)

	productID := uuid.New()

	mockDB.On("First", mock.Anything, "id = ? and deleted_at is null", productID).Return(withErr)

	result, err := repo.GetByID(productID)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}
