package cart

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/errutil"
)

type Service struct {
	CartRepository         ICartRepository
	CartProductsRepository ICartProductRepository
}

func NewService(cartRepository ICartRepository, cartProductsRepository ICartProductRepository) *Service {
	return &Service{
		CartRepository:         cartRepository,
		CartProductsRepository: cartProductsRepository,
	}
}

func (s *Service) LoadCart(clientID uuid.UUID) (*Cart, error) {
	cart, err := s.CartRepository.Get(clientID)
	if err != nil {
		if !errors.Is(err, errutil.ErrRecordNotFound) {
			return nil, err
		}

		cart, err = s.CartRepository.Create(clientID)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}

func (s *Service) GetFullCart(clientID uuid.UUID) (*Cart, error) {
	cart, err := s.CartRepository.Get(clientID)
	if err != nil {
		return nil, err
	}

	products, err := s.CartProductsRepository.GetByCartID(context.Background(), cart.ID)
	if err != nil {
		return nil, err
	}

	cart.Products = products

	return cart, nil
}

func (s *Service) Cleanup(clientID uuid.UUID) error {
	cart, err := s.LoadCart(clientID)
	if err != nil {
		return err
	}

	products, err := s.CartProductsRepository.GetByCartID(context.Background(), cart.ID)
	if err != nil {
		return err
	}

	for _, p := range products {
		err = s.CartProductsRepository.DeleteByProductID(context.Background(), cart.ID, p.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) AddProduct(ctx context.Context, clientID uuid.UUID, product *Product) error {
	cart, err := s.LoadCart(clientID)
	if err != nil {
		return err
	}

	return s.CartProductsRepository.Create(ctx, cart.ID, product)
}

func (s *Service) RemoveProduct(ctx context.Context, clientID uuid.UUID, productID uuid.UUID) error {
	cart, err := s.LoadCart(clientID)
	if err != nil {
		return err
	}

	return s.CartProductsRepository.DeleteByProductID(ctx, cart.ID, productID)
}

func (s *Service) EditProduct(ctx context.Context, clientID uuid.UUID, product *Product) error {
	cart, err := s.LoadCart(clientID)
	if err != nil {
		return err
	}

	return s.CartProductsRepository.UpdateProductByProductID(ctx, cart.ID, product.ProductID, product)
}
