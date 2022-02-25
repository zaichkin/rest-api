package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/julienschmidt/httprouter"
	"market/internal/domain/product"
	"market/pkg/database/postgresql"
)

type ProductStorage struct {
	client postgresql.Connect
}

func (p *ProductStorage) CreateRowDB(ctx context.Context, dto product.CreateProductDTO) (product.Product, error) {
	sql := `
	INSERT INTO product(title, price, size, category, gender, brand, descript, img) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, title, price, size, category, gender, brand, descript, img
	`

	var elem product.Product
	err := p.client.QueryRow(ctx, sql, dto.Title, dto.Price, dto.Size, dto.Category, dto.Gender, dto.Brand, dto.Description, dto.Img).Scan(&elem.Id, &elem.Title, &elem.Price, &elem.Size, &elem.Category, &elem.Gender, &elem.Brand, &elem.Description, &elem.Img)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return product.Product{}, err
	}
	return elem, nil
}

func (p *ProductStorage) DeleteRowDB(ctx context.Context, dto product.DeleteProductDTO) error {
	sql := `
	DELETE FROM product 
	WHERE id = $1
	`
	_, err := p.client.Exec(ctx, sql, dto.Id)
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

func (p *ProductStorage) GetRowDB(ctx context.Context, dto product.GetProductDTO) (product.Product, error) {
	sql := `
	SELECT id, title, price, size, category, gender, brand, descript, img 
	FROM product
	WHERE id = $1;
	`
	var elem product.Product
	if err := p.client.QueryRow(ctx, sql, dto.Id).Scan(&elem.Id, &elem.Title, &elem.Price, &elem.Size, &elem.Category, &elem.Gender, &elem.Brand, &elem.Description, &elem.Img); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return product.Product{}, err
	}
	return elem, nil
}

func (p *ProductStorage) UpdateRowDB(ctx context.Context, dto product.UpdateProductDTO) (product.Product, error) {
	sql := `
	SELECT id, title, price, size, category, gender, brand, descript, img 
	FROM product
	WHERE id = $1;
	`
	var oldelem product.Product
	if err := p.client.QueryRow(ctx, sql, dto.Id).Scan(&oldelem.Id, &oldelem.Title, &oldelem.Price, &oldelem.Size, &oldelem.Category, &oldelem.Gender, &oldelem.Brand, &oldelem.Description, &oldelem.Img); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return product.Product{}, err
	}

	if dto.Title == nil {
		dto.Title = &oldelem.Title
	}
	if dto.Price == nil {
		dto.Price = &oldelem.Price
	}
	if dto.Size == nil {
		dto.Size = &oldelem.Size
	}
	if dto.Category == nil {
		dto.Category = &oldelem.Category
	}
	if dto.Gender == nil {
		dto.Gender = &oldelem.Gender
	}
	if dto.Brand == nil {
		dto.Brand = &oldelem.Brand
	}
	if dto.Description == nil {
		dto.Description = &oldelem.Description
	}
	if dto.Img == nil {
		dto.Img = &oldelem.Img
	}

	sql = `
	UPDATE product
	SET title = $1,
		price = $2, 
		size = $3, 
		category = $4, 
		gender = $5, 
		brand = $6, 
		descript = $7, 
		img = $8
	WHERE id = $9
	RETURNING id, title, price, size, category, gender, brand, descript, img
	`
	var elem product.Product
	err := p.client.QueryRow(ctx, sql, dto.Title, dto.Price, dto.Size, dto.Category, dto.Gender, dto.Brand, dto.Description, dto.Img, dto.Id).Scan(&elem.Id, &elem.Title, &elem.Price, &elem.Size, &elem.Category, &elem.Gender, &elem.Brand, &elem.Description, &elem.Img)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			err = err.(*pgconn.PgError)
			fmt.Printf("Code: %s, Message: %s, Detail: %s, ", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return product.Product{}, err
	}

	return elem, nil
}

func (p *ProductStorage) AllRowsDB(ctx context.Context, params httprouter.Params) ([]product.Product, error) {
	sql := `
	SELECT id, title, price, size, category, gender, brand, descript, img  
	FROM product
	` + WorkWithSelect(params)
	var elem product.Product
	var list []product.Product
	rows, err := p.client.Query(ctx, sql)
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
		if err := rows.Scan(&elem.Id, &elem.Title, &elem.Price, &elem.Size, &elem.Category, &elem.Gender, &elem.Brand, &elem.Description, &elem.Img); err != nil {
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

func WorkWithSelect(params httprouter.Params) string {
	sort := params.ByName("sort") // priceh pricel newlist
	//where := params.ByName("brand")
	other := ""
	switch sort {
	case "priceh":
		other += " ORDER BY price DESC "
	case "pricel":
		other += " ORDER BY price "
	case "newlist":
		other += " OREDER BY newlist "
	}
	return other
}

func NewProduct(conn postgresql.Connect) *ProductStorage {
	return &ProductStorage{
		client: conn,
	}
}
