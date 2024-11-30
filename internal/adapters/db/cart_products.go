package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/cart"
)

type PostgresCartProductsRepository struct {
	db IDB
}

type CartProductsPostgres struct {
	BaseModel
	CartID    uuid.UUID       `gorm:"type:uuid"`
	ProductID uuid.UUID       `gorm:"type:uuid"`
	Quantity  int             `gorm:"quantity"`
	Comments  string          `gorm:"comments"`
	Cart      CartPostgres    `gorm:"foreignKey:CartID"`
	Product   ProductPostgres `gorm:"foreignKey:ProductID"`
}

func (op *CartProductsPostgres) TableName() string {
	return "cart_products"
}

func NewPostgresCartProductsRepository(db IDB) cart.ICartProductRepository {
	return &PostgresCartProductsRepository{db: db}
}

func (p *PostgresCartProductsRepository) Create(_ context.Context, cartID uuid.UUID, product *cart.Product) error {
	cartProduct := CartProductsPostgres{
		CartID:    cartID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
		Comments:  product.Comments,
	}

	result := p.db.Create(&cartProduct)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PostgresCartProductsRepository) GetByCartID(_ context.Context, cartID uuid.UUID) ([]*cart.Product, error) {
	var cartProducts []CartProductsPostgres
	err := p.db.Find(&cartProducts, "cart_id = ?", cartID).Error
	if err != nil {
		return nil, err
	}

	var products []*cart.Product
	for _, cp := range cartProducts {
		products = append(products, &cart.Product{
			ProductID: cp.ProductID,
			Quantity:  cp.Quantity,
			Comments:  cp.Comments,
		})
	}

	return products, nil
}

func (p *PostgresCartProductsRepository) DeleteByProductID(_ context.Context, cartID, productID uuid.UUID) error {
	return p.db.Delete(&CartProductsPostgres{}, "cart_id = ? AND product_id = ?", cartID, productID).Error
}

func (p *PostgresCartProductsRepository) UpdateProductByProductID(_ context.Context, cartID, productID uuid.UUID, product *cart.Product) error {
	return p.db.Model(&CartProductsPostgres{}).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		Updates(map[string]interface{}{
			"quantity": product.Quantity,
			"comments": product.Comments,
		}).Error
}
