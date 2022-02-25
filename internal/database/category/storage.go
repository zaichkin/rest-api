package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"market/internal/domain/category"
	"market/pkg/database/postgresql"
)

type CategoryStorage struct {
	client postgresql.Connect
}

func (c *CategoryStorage) CreateRowDB(ctx context.Context, dto category.CreateCategoryDTO) (category.Category, error) {
	sql := `
	INSERT INTO category(parent, name) 
	VALUES ($1, $2)
	RETURNING id, parent, name
	`

	var elem category.Category
	err := c.client.QueryRow(ctx, sql, dto.Parent, dto.Name).Scan(&elem.Id, &elem.Parent, &elem.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return category.Category{}, err
	}
	return elem, nil
}

func (c *CategoryStorage) DeleteRowDB(ctx context.Context, dto category.DeleteCategoryDTO) error {
	sql := `
	DELETE FROM category
	WHERE id = $1
	`
	_, err := c.client.Exec(ctx, sql, dto.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError) //for logger
		}
		return err
	}
	return nil
}

func (c *CategoryStorage) GetRowDB(ctx context.Context, dto category.GetCategoryDTO) (category.Category, error) {
	sql := `
	SELECT id, parent, name
	FROM category
	WHERE id = $1;
	`
	var elem category.Category
	if err := c.client.QueryRow(ctx, sql, dto.Id).Scan(&elem.Id, &elem.Parent, &elem.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return category.Category{}, err
	}

	return elem, nil
}

func (c *CategoryStorage) UpdateRowDB(ctx context.Context, dto category.UpdateCategoryDTO) (category.Category, error) {
	sql := `
	SELECT id, parent, name
	FROM category
	WHERE id = $1;
	`
	var oldelem category.Category
	if err := c.client.QueryRow(ctx, sql, dto.Id).Scan(&oldelem.Id, &oldelem.Parent, &oldelem.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return category.Category{}, err
	}

	if dto.Parent == nil {
		dto.Parent = &oldelem.Parent
	}
	if dto.Name == nil {
		dto.Name = &oldelem.Name
	}

	sql = `
	UPDATE category
	SET parent = $1,
		name = $2
	WHERE id = $3
	RETURNING id, parent, name
	`
	var elem category.Category
	err := c.client.QueryRow(ctx, sql, dto.Parent, dto.Name, dto.Id).Scan(&elem.Id, &elem.Parent, &elem.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return category.Category{}, err
	}
	return elem, nil
}

func (c *CategoryStorage) AllRowsDB(ctx context.Context) ([]category.Category, error) {
	sql := `
	SELECT id, parent, name 
	FROM category;
	`
	var elem category.Category
	var list []category.Category
	rows, err := c.client.Query(ctx, sql)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&elem.Id, &elem.Parent, &elem.Name); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				err = err.(*pgconn.PgError)
			}
			return nil, err
		}
		list = append(list, elem)
	}

	return list, nil
}

func NewCategory(conn postgresql.Connect) *CategoryStorage {
	return &CategoryStorage{
		client: conn,
	}
}
