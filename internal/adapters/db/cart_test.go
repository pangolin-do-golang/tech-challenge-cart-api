package db_test

import (
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/errutil"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

var noErr = &gorm.DB{}
var withErr = &gorm.DB{Error: errutil.ErrRecordNotFound}

func TestCreateCart_Success(t *testing.T) {
	mockDB := new(mocks.IDB)
	mockDB.On("Create", mock.Anything).Return(noErr)
	repo := db.NewPostgresCartRepository(mockDB)

	clientID := uuid.New()

	cart, err := repo.Create(clientID)

	assert.NoError(t, err)
	assert.NotNil(t, cart)
	assert.Equal(t, clientID, cart.ClientID)
	mockDB.AssertExpectations(t)
}

func TestCreateCart_DBError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartRepository(mockDB)

	clientID := uuid.New()
	mockDB.On("Create", mock.Anything).Return(withErr)

	cart, err := repo.Create(clientID)

	assert.Error(t, err)
	assert.Nil(t, cart)
	mockDB.AssertExpectations(t)
}

func TestGetCart_Success(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartRepository(mockDB)

	clientID := uuid.New()
	expectedCart := &db.CartPostgres{ClientID: clientID}
	mockDB.On("First", mock.Anything, "client_id = ?", clientID).Return(noErr).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*db.CartPostgres)
		*arg = *expectedCart
	})

	cart, err := repo.Get(clientID)

	assert.NoError(t, err)
	assert.NotNil(t, cart)
	assert.Equal(t, clientID, cart.ClientID)
	mockDB.AssertExpectations(t)
}

func TestGetCart_NotFound(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartRepository(mockDB)

	clientID := uuid.New()
	mockDB.On("First", mock.Anything, "client_id = ?", clientID).Return(withErr)

	cart, err := repo.Get(clientID)

	assert.Error(t, err)
	assert.Nil(t, cart)
	assert.Equal(t, errutil.ErrRecordNotFound, err)
	mockDB.AssertExpectations(t)
}

func TestGetCart_DBError(t *testing.T) {
	mockDB := new(mocks.IDB)
	repo := db.NewPostgresCartRepository(mockDB)

	clientID := uuid.New()
	mockDB.On("First", mock.Anything, "client_id = ?", clientID).Return(withErr)

	cart, err := repo.Get(clientID)

	assert.Error(t, err)
	assert.Nil(t, cart)
	mockDB.AssertExpectations(t)
}
