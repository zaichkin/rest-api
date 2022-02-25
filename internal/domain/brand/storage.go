package brand

import (
	"context"
)

type Storage interface {
	CreateRowDB(ctx context.Context, md CreateBrandDTO) (Brand, error)
	DeleteRowDB(ctx context.Context, md DeleteBrandDTO) error
	GetRowDB(ctx context.Context, md GetBrandDTO) (Brand, error)
	UpdateRowDB(ctx context.Context, md UpdateBrandDTO) (Brand, error)
	AllRowsDB(ctx context.Context) ([]Brand, error)
}
