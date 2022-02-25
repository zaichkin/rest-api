package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/julienschmidt/httprouter"
	"market/internal/domain/worker"
	"market/pkg/database/postgresql"
)

type WorkerStorage struct {
	client postgresql.Connect
}

func (w *WorkerStorage) CreateRowDB(ctx context.Context, dto worker.CreateWorkerDTO) (worker.Worker, error) {
	sql := `
	INSERT INTO worker(name, descript, workspace) 
	VALUES ($1, $2, $3)
	RETURNING id, name, descript, workspace
	`
	var elem worker.Worker
	err := w.client.QueryRow(ctx, sql, dto.Name, dto.Description, dto.Workspace).Scan(&elem.Id, &elem.Name, &elem.Description, &elem.Workspace)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return worker.Worker{}, err
	}
	return elem, nil
}

func (w *WorkerStorage) DeleteRowDB(ctx context.Context, dto worker.DeleteWorkerDTO) error {
	sql := `
	DELETE FROM worker
	WHERE id = $1
	`
	_, err := w.client.Exec(ctx, sql, dto.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError) //for logger
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return err
	}
	return nil
}

func (w *WorkerStorage) GetRowDB(ctx context.Context, dto worker.GetWorkerDTO) (worker.Worker, error) {
	sql := `
	SELECT id, name, descript, workspace 
	FROM worker
	WHERE id = $1;
	`
	var elem worker.Worker
	if err := w.client.QueryRow(ctx, sql, dto.Id).Scan(&elem.Id, &elem.Name, &elem.Description, &elem.Workspace); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return worker.Worker{}, err
	}

	return elem, nil
}

func (w *WorkerStorage) UpdateRowDB(ctx context.Context, dto worker.UpdateWorkerDTO) (worker.Worker, error) {
	sql := `
	SELECT id, name, descript, workspace 
	FROM worker
	WHERE id = $1;
	`
	var oldelem worker.Worker
	if err := w.client.QueryRow(ctx, sql, dto.Id).Scan(&oldelem.Id, &oldelem.Name, &oldelem.Description, &oldelem.Workspace); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return worker.Worker{}, err
	}

	if dto.Name == nil {
		dto.Name = &oldelem.Name
	}
	if dto.Description == nil {
		dto.Description = &oldelem.Description
	}
	if dto.Workspace == nil {
		dto.Workspace = &oldelem.Workspace
	}

	sql = `
	UPDATE worker
	SET name = $1,
		descript = $2,
		workspace = $3
	WHERE id = $4
	RETURNING id, name, descript, workspace
	`
	var elem worker.Worker
	err := w.client.QueryRow(ctx, sql, dto.Name, dto.Description, dto.Workspace, dto.Id).Scan(&elem.Id, &elem.Name, &elem.Description, &elem.Workspace)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return worker.Worker{}, err
	}

	return elem, nil
}

func (w *WorkerStorage) AllRowsDB(ctx context.Context, params httprouter.Params) ([]worker.Worker, error) {
	sql := `
	SELECT id, name, descript, workspace 
	FROM worker`
	fmt.Println(sql + Filter(params))
	var elem worker.Worker
	var list []worker.Worker
	rows, err := w.client.Query(ctx, sql)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&elem.Id, &elem.Name, &elem.Description, &elem.Workspace); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				err = err.(*pgconn.PgError)
				fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
			}
			return nil, err
		}
		list = append(list, elem)
	}

	return list, nil
}

func Filter(param httprouter.Params) string {
	parent := param.ByName("parent")
	if parent == "" {
		return ";"
	}
	return "WHERE parent = " + parent
}

func NewWorker(conn postgresql.Connect) *WorkerStorage {
	return &WorkerStorage{
		client: conn,
	}
}
