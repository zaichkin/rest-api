package brand

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"

	"market/internal/domain/brand"
	"market/pkg/database/postgresql"
)

type BrandStorage struct {
	client postgresql.Connect
}

func (b *BrandStorage) CreateRowDB(ctx context.Context, dto brand.CreateBrandDTO) (brand.Brand, error) {
	sql := `
	INSERT INTO brand(title, img, descript) 
	VALUES ($1, $2, $3) 
	RETURNING id, title, img, descript
	`
	var elem brand.Brand
	err := b.client.QueryRow(ctx, sql, dto.Title, dto.Img, dto.Description).Scan(&elem.Id, &elem.Title, &elem.Img, &elem.Description)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return brand.Brand{}, err
	}
	return elem, nil
}

func (b *BrandStorage) DeleteRowDB(ctx context.Context, dto brand.DeleteBrandDTO) error {
	sql := `
	DELETE FROM brand
	WHERE id = $1
	`
	_, err := b.client.Exec(ctx, sql, dto.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError) //for logger
		}
		return err
	}
	return nil
}

func (b *BrandStorage) GetRowDB(ctx context.Context, dto brand.GetBrandDTO) (brand.Brand, error) {
	sql := `
	SELECT id, title, img, descript
	FROM brand
	WHERE id = $1;
	`
	var elem brand.Brand
	if err := b.client.QueryRow(ctx, sql, dto.Id).Scan(&elem.Id, &elem.Title, &elem.Img, &elem.Description); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return brand.Brand{}, err
	}
	return elem, nil
}

func (b *BrandStorage) UpdateRowDB(ctx context.Context, dto brand.UpdateBrandDTO) (brand.Brand, error) {
	sql := `
	SELECT id, title, img, descript
	FROM brand
	WHERE id = $1;
	`
	var oldelem brand.Brand
	if err := b.client.QueryRow(ctx, sql, dto.Id).Scan(&oldelem.Id, &oldelem.Title, &oldelem.Img, &oldelem.Description); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return brand.Brand{}, err
	}

	if dto.Title == nil {
		dto.Title = &oldelem.Title
	}
	if dto.Description == nil {
		dto.Description = &oldelem.Description
	}
	if dto.Img == nil {
		dto.Img = &oldelem.Img
	}

	sql = `
	UPDATE brand
	SET title = $1,
		descript = $2,
		img = $3
	WHERE id = $4
	RETURNING id, title, descript, img
	`
	var elem brand.Brand
	err := b.client.QueryRow(ctx, sql, dto.Title, dto.Description, dto.Img, dto.Id).Scan(&elem.Id, &elem.Title, &elem.Description, &elem.Img)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return brand.Brand{}, err
	}
	return elem, nil
}

func (b *BrandStorage) AllRowsDB(ctx context.Context) ([]brand.Brand, error) {
	sql := `
	SELECT id, title, img, descript
	FROM brand;
	`
	var elem brand.Brand
	var list []brand.Brand
	rows, err := b.client.Query(ctx, sql)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&elem.Id, &elem.Title, &elem.Img, &elem.Description); err != nil {
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

func NewBrand(conn postgresql.Connect) *BrandStorage {
	return &BrandStorage{
		client: conn,
	}
}
