package db_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCartProduct_Success(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	product := &cart.Product{
		ProductID: uuid.New(),
		Quantity:  1,
		Comments:  "Test comment",
	}

	mockDB.On("Create", mock.Anything).Return(noErr)

	err := repo.Create(context.Background(), cartID, product)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestCreateCartProduct_DBError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	product := &cart.Product{
		ProductID: uuid.New(),
		Quantity:  1,
		Comments:  "Test comment",
	}

	mockDB.On("Create", mock.Anything).Return(withErr)

	err := repo.Create(context.Background(), cartID, product)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

func TestGetCartProductsByCartID_Success(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	expectedProducts := []db.CartProductsPostgres{
		{CartID: cartID, ProductID: uuid.New(), Quantity: 1, Comments: "Test comment"},
	}

	mockDB.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(noErr).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]db.CartProductsPostgres)
		*arg = expectedProducts
	})

	products, err := repo.GetByCartID(context.Background(), cartID)

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, expectedProducts[0].ProductID, products[0].ProductID)
	mockDB.AssertExpectations(t)
}

func TestGetCartProductsByCartID_DBError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()

	mockDB.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(withErr)

	products, err := repo.GetByCartID(context.Background(), cartID)

	assert.Error(t, err)
	assert.Nil(t, products)
	mockDB.AssertExpectations(t)
}

func TestDeleteCartProductByProductID_Success(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	productID := uuid.New()

	mockDB.On("Delete", mock.Anything, "cart_id = ? AND product_id = ?", cartID, productID).Return(noErr)

	err := repo.DeleteByProductID(context.Background(), cartID, productID)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteCartProductByProductID_DBError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	productID := uuid.New()

	mockDB.On("Delete", mock.Anything, "cart_id = ? AND product_id = ?", cartID, productID).Return(withErr)

	err := repo.DeleteByProductID(context.Background(), cartID, productID)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

/*
func TestUpdateCartProductByProductID_Success(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	productID := uuid.New()
	product := &cart.Product{
		ProductID: productID,
		Quantity:  2,
		Comments:  "Updated comment",
	}

	mockDB.On("Model", mock.Anything).Return(mockDB)
	mockDB.On("Where", "cart_id = ? AND product_id = ?", cartID, productID).Return(mockDB)
	mockDB.On("Updates", mock.Anything).Return(noErr)

	err := repo.UpdateProductByProductID(context.Background(), cartID, productID, product)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateCartProductByProductID_DBError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartProductsRepository(mockDB)

	cartID := uuid.New()
	productID := uuid.New()
	product := &cart.Product{
		ProductID: productID,
		Quantity:  2,
		Comments:  "Updated comment",
	}

	mockDB.On("Model", mock.Anything).Return(mockDB)
	mockDB.On("Where", "cart_id = ? AND product_id = ?", cartID, productID).Return(mockDB)
	mockDB.On("Updates", mock.Anything).Return(withErr)

	err := repo.UpdateProductByProductID(context.Background(), cartID, productID, product)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}
*/
