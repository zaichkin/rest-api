package product

import (
	"context"
	"github.com/julienschmidt/httprouter"
)

type Storage interface {
	CreateRowDB(ctx context.Context, dto CreateProductDTO) (Product, error)
	DeleteRowDB(ctx context.Context, dto DeleteProductDTO) error
	GetRowDB(ctx context.Context, dto GetProductDTO) (Product, error)
	UpdateRowDB(ctx context.Context, dto UpdateProductDTO) (Product, error)
	AllRowsDB(ctx context.Context, params httprouter.Params) ([]Product, error)
}
