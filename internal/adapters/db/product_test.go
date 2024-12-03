package db_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/product"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func TestProductSearch(t *testing.T) {
	d, m, err := sqlmock.New()
	conn, err := gorm.Open(postgres.New(postgres.Config{Conn: d, DriverName: "postgres"}))

	m.ExpectQuery("SELECT .+").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "description", "category", "price"}).
			AddRow(uuid.New(), "Test Product", "Test Description", "Test Category", 100.0).
			AddRow(uuid.New(), "Test Product 2", "Test Description 2", "Test Category", 200.0),
	)
	repo := db.NewPostgresProductRepository(conn)
	_, err = repo.Search("search", "category")

	assert.NoError(t, err)
}

func TestProductDelete(t *testing.T) {
	d, m, err := sqlmock.New()
	conn, err := gorm.Open(postgres.New(postgres.Config{Conn: d, DriverName: "postgres"}))

	m.ExpectQuery("SELECT .+").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "description", "category", "price"}).
			AddRow(uuid.New(), "Test Product", "Test Description", "Test Category", 100.0),
	)

	m.ExpectBegin()
	m.ExpectExec("UPDATE .+").WillReturnResult(sqlmock.NewResult(0, 1))
	m.ExpectCommit()
	repo := db.NewPostgresProductRepository(conn)
	err = repo.Delete(uuid.New())

	assert.NoError(t, err)
}
