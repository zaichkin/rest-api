package worker

import (
	"context"
	"github.com/julienschmidt/httprouter"
)

type Storage interface {
	CreateRowDB(ctx context.Context, dto CreateWorkerDTO) (Worker, error)
	DeleteRowDB(ctx context.Context, dto DeleteWorkerDTO) error
	GetRowDB(ctx context.Context, dto GetWorkerDTO) (Worker, error)
	UpdateRowDB(ctx context.Context, dto UpdateWorkerDTO) (Worker, error)
	AllRowsDB(ctx context.Context, param httprouter.Params) ([]Worker, error)
}
