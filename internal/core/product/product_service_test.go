package product_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/product"
	"github.com/pangolin-do-golang/tech-challenge-cart-api/mocks"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	products := &[]product.Product{
		{
			Id:          uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc"),
			Name:        "Teste",
			Description: "Teste Desc",
			Category:    "teste",
			Price:       120,
		},
		{
			Id:          uuid.MustParse("518a84c1-4199-48a0-8ac3-ebcf54ab783e"),
			Name:        "Teste 2",
			Description: "Teste Desc 2",
			Category:    "teste",
			Price:       150,
		},
	}

	type fields struct {
		genRepository func() product.Repository
	}
	type args struct {
		search   string
		category string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]product.Product
		wantErr bool
	}{
		{
			name: "returns error from search method",
			fields: fields{
				genRepository: func() product.Repository {
					m := new(mocks.Repository)
					m.On("Search", "search", "category").Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				search:   "search",
				category: "category",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "successful run",
			fields: fields{
				genRepository: func() product.Repository {
					m := new(mocks.Repository)
					m.On("Search", "search", "category").Return(products, nil)
					return m
				},
			},
			args: args{
				search:   "search",
				category: "category",
			},
			want:    products,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := product.NewProductService(tt.fields.genRepository())
			got, err := s.Search(tt.args.search, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	type fields struct {
		genRepository func() product.Repository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns error from delete method",
			fields: fields{
				genRepository: func() product.Repository {
					m := new(mocks.Repository)
					m.On("Delete", uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc")).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				id: uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc"),
			},
			wantErr: true,
		},
		{
			name: "successful run",
			fields: fields{
				genRepository: func() product.Repository {
					m := new(mocks.Repository)
					m.On("Delete", uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc")).Return(nil)
					return m
				},
			},
			args: args{
				id: uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := product.NewProductService(tt.fields.genRepository())
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetByID(t *testing.T) {
	p := &product.Product{
		Id: uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc"),
	}

	type fields struct {
		genRepository func() product.Repository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *product.Product
		wantErr bool
	}{
		{
			name: "returns error from get by id method",
			fields: fields{
				genRepository: func() product.Repository {
					m := new(mocks.Repository)
					m.On("GetByID", uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc")).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				id: uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "successful run",
			fields: fields{
				genRepository: func() product.Repository {
					m := new(mocks.Repository)
					m.On("GetByID", uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc")).Return(p, nil)
					return m
				},
			},
			args: args{
				id: uuid.MustParse("0c88073f-b024-4d50-9235-b5ac53e887bc"),
			},
			want:    p,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := product.NewProductService(tt.fields.genRepository())
			got, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
