package category

import (
	"context"
)

type Storage interface {
	CreateRowDB(ctx context.Context, dto CreateCategoryDTO) (Category, error)
	DeleteRowDB(ctx context.Context, dto DeleteCategoryDTO) error
	GetRowDB(ctx context.Context, dto GetCategoryDTO) (Category, error)
	UpdateRowDB(ctx context.Context, dto UpdateCategoryDTO) (Category, error)
	AllRowsDB(ctx context.Context) ([]Category, error)
}
