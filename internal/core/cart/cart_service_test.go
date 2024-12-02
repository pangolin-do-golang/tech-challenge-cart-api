package cart_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/errutil"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestService_LoadCart(t *testing.T) {
	var (
		clientID = uuid.MustParse("e6fee820-9d3e-4ec2-92be-0fdfa6073f58")
		c        = &cart.Cart{
			ID:       uuid.MustParse("23636e91-a39a-4098-8614-bf600dcadc6e"),
			ClientID: clientID,
		}
	)

	type fields struct {
		genCartRepository func() cart.ICartRepository
	}
	type args struct {
		clientID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cart.Cart
		wantErr bool
	}{
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", clientID).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creates new cart",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", clientID).Return(nil, errutil.ErrRecordNotFound)
					m.On("Create", clientID).Return(c, nil)
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			want:    c,
			wantErr: false,
		},
		{
			name: "returns existing cart",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", clientID).Return(c, nil)
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			want:    c,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cart.Service{
				CartRepository: tt.fields.genCartRepository(),
				CartProductsRepository: func() cart.ICartProductRepository {
					return new(mocks.ICartProductRepository)
				}(),
			}
			got, err := s.LoadCart(tt.args.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetFullCart(t *testing.T) {
	var (
		clientID = uuid.MustParse("e6fee820-9d3e-4ec2-92be-0fdfa6073f58")
		c        = &cart.Cart{
			ID:       uuid.MustParse("23636e91-a39a-4098-8614-bf600dcadc6e"),
			ClientID: clientID,
		}
		p = []*cart.Product{
			{
				ProductID: uuid.MustParse("da311633-ebaf-4076-b119-79ea5eabcaaa"),
				Quantity:  1,
				Comments:  "",
			},
		}
	)

	fullCart := c
	c.Products = p

	type fields struct {
		genCartRepository         func() cart.ICartRepository
		genCartProductsRepository func() cart.ICartProductRepository
	}
	type args struct {
		clientID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cart.Cart
		wantErr bool
	}{
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", clientID).Return(nil, errors.New("error"))
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					return new(mocks.ICartProductRepository)
				},
			},
			args: args{
				clientID: clientID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error from cart product repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", clientID).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("GetByCartID", mock.Anything, c.ID).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns cart with products",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", clientID).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("GetByCartID", mock.Anything, c.ID).Return(p, nil)
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			want:    fullCart,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cart.Service{
				CartRepository:         tt.fields.genCartRepository(),
				CartProductsRepository: tt.fields.genCartProductsRepository(),
			}
			got, err := s.GetFullCart(tt.args.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFullCart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Cleanup(t *testing.T) {
	var (
		clientID = uuid.MustParse("e6fee820-9d3e-4ec2-92be-0fdfa6073f58")
		c        = &cart.Cart{
			ID:       uuid.MustParse("23636e91-a39a-4098-8614-bf600dcadc6e"),
			ClientID: clientID,
		}
		p = []*cart.Product{
			{
				ProductID: uuid.MustParse("da311633-ebaf-4076-b119-79ea5eabcaaa"),
				Quantity:  1,
				Comments:  "",
			},
		}
	)

	type fields struct {
		genCartRepository         func() cart.ICartRepository
		genCartProductsRepository func() cart.ICartProductRepository
	}
	type args struct {
		clientID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(nil, errors.New("error"))
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					return new(mocks.ICartProductRepository)
				},
			},
			args: args{
				clientID: clientID,
			},
			wantErr: true,
		},
		{
			name: "returns error from cart products repository get method",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("GetByCartID", mock.Anything, c.ID).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			wantErr: true,
		},
		{
			name: "returns error from cart products repository delete method",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("GetByCartID", mock.Anything, c.ID).Return(p, nil)
					m.On("DeleteByProductID", mock.Anything, c.ID, p[0].ProductID).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			wantErr: true,
		},
		{
			name: "returns no error from successful run",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("GetByCartID", mock.Anything, c.ID).Return(p, nil)
					m.On("DeleteByProductID", mock.Anything, c.ID, p[0].ProductID).Return(nil)
					return m
				},
			},
			args: args{
				clientID: clientID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cart.Service{
				CartRepository:         tt.fields.genCartRepository(),
				CartProductsRepository: tt.fields.genCartProductsRepository(),
			}
			if err := s.Cleanup(tt.args.clientID); (err != nil) != tt.wantErr {
				t.Errorf("Cleanup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_AddProduct(t *testing.T) {
	var (
		clientID = uuid.MustParse("e6fee820-9d3e-4ec2-92be-0fdfa6073f58")
		c        = &cart.Cart{
			ID:       uuid.MustParse("23636e91-a39a-4098-8614-bf600dcadc6e"),
			ClientID: clientID,
		}
		p = &cart.Product{
			ProductID: uuid.MustParse("da311633-ebaf-4076-b119-79ea5eabcaaa"),
			Quantity:  1,
			Comments:  "",
		}
	)

	type fields struct {
		genCartRepository         func() cart.ICartRepository
		genCartProductsRepository func() cart.ICartProductRepository
	}
	type args struct {
		ctx      context.Context
		clientID uuid.UUID
		product  *cart.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(nil, errors.New("error"))
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					return new(mocks.ICartProductRepository)
				},
			},
			args: args{
				clientID: clientID,
			},
			wantErr: true,
		},
		{
			name: "returns error from cart products repository create method",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("Create", mock.Anything, c.ID, p).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				clientID: clientID,
				product:  p,
			},
			wantErr: true,
		},
		{
			name: "successfully creates product",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("Create", mock.Anything, c.ID, p).Return(nil)
					return m
				},
			},
			args: args{
				clientID: clientID,
				product:  p,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cart.Service{
				CartRepository:         tt.fields.genCartRepository(),
				CartProductsRepository: tt.fields.genCartProductsRepository(),
			}
			if err := s.AddProduct(context.TODO(), tt.args.clientID, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("AddProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RemoveProduct(t *testing.T) {
	var (
		clientID = uuid.MustParse("e6fee820-9d3e-4ec2-92be-0fdfa6073f58")
		c        = &cart.Cart{
			ID:       uuid.MustParse("23636e91-a39a-4098-8614-bf600dcadc6e"),
			ClientID: clientID,
		}
		p = &cart.Product{
			ProductID: uuid.MustParse("da311633-ebaf-4076-b119-79ea5eabcaaa"),
			Quantity:  1,
			Comments:  "",
		}
	)

	type fields struct {
		genCartRepository         func() cart.ICartRepository
		genCartProductsRepository func() cart.ICartProductRepository
	}
	type args struct {
		ctx       context.Context
		clientID  uuid.UUID
		productID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(nil, errors.New("error"))
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					return new(mocks.ICartProductRepository)
				},
			},
			args: args{
				clientID:  clientID,
				productID: p.ProductID,
			},
			wantErr: true,
		},
		{
			name: "returns error from cart products repository delete method",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("DeleteByProductID", mock.Anything, c.ID, p.ProductID).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				clientID:  clientID,
				productID: p.ProductID,
			},
			wantErr: true,
		},
		{
			name: "returns no error for successful run",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("DeleteByProductID", mock.Anything, c.ID, p.ProductID).Return(nil)
					return m
				},
			},
			args: args{
				clientID:  clientID,
				productID: p.ProductID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cart.Service{
				CartRepository:         tt.fields.genCartRepository(),
				CartProductsRepository: tt.fields.genCartProductsRepository(),
			}
			if err := s.RemoveProduct(context.TODO(), tt.args.clientID, tt.args.productID); (err != nil) != tt.wantErr {
				t.Errorf("RemoveProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_EditProduct(t *testing.T) {
	var (
		clientID = uuid.MustParse("e6fee820-9d3e-4ec2-92be-0fdfa6073f58")
		c        = &cart.Cart{
			ID:       uuid.MustParse("23636e91-a39a-4098-8614-bf600dcadc6e"),
			ClientID: clientID,
		}
		p = &cart.Product{
			ProductID: uuid.MustParse("da311633-ebaf-4076-b119-79ea5eabcaaa"),
			Quantity:  1,
			Comments:  "",
		}
	)

	type fields struct {
		genCartRepository         func() cart.ICartRepository
		genCartProductsRepository func() cart.ICartProductRepository
	}
	type args struct {
		ctx      context.Context
		clientID uuid.UUID
		product  *cart.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(nil, errors.New("error"))
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					return new(mocks.ICartProductRepository)
				},
			},
			args: args{
				clientID: clientID,
				product:  p,
			},
			wantErr: true,
		},
		{
			name: "returns error from cart repository",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("UpdateProductByProductID", mock.Anything, c.ID, p.ProductID, p).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				clientID: clientID,
				product:  p,
			},
			wantErr: true,
		},
		{
			name: "successful run",
			fields: fields{
				genCartRepository: func() cart.ICartRepository {
					m := new(mocks.ICartRepository)
					m.On("Get", mock.Anything).Return(c, nil)
					return m
				},
				genCartProductsRepository: func() cart.ICartProductRepository {
					m := new(mocks.ICartProductRepository)
					m.On("UpdateProductByProductID", mock.Anything, c.ID, p.ProductID, p).Return(nil)
					return m
				},
			},
			args: args{
				clientID: clientID,
				product:  p,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cart.Service{
				CartRepository:         tt.fields.genCartRepository(),
				CartProductsRepository: tt.fields.genCartProductsRepository(),
			}
			if err := s.EditProduct(context.TODO(), tt.args.clientID, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("EditProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
