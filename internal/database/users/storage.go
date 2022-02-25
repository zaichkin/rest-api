package users

import (
	"context"
	"market/internal/domain/users"
	"market/pkg/database/postgresql"
)

type UserStorage struct {
	client postgresql.Connect
}

func (u *UserStorage) CreateUser(ctx context.Context, dto users.SingUpDTO) error {
	sql := `
	INSERT INTO users (uuid, login, password) 
	VALUES ($1, $2, $3)
	RETURNING uuid
	`

	if err := u.client.QueryRow(ctx, sql).Scan(); err != nil {
		return err
	}
	return nil
}

func (u *UserStorage) GetUser(ctx context.Context, dto users.SingInDTO) error {
	sql := `
	SELECT uuid, login, password
	FROM users 
	WHERE login = $1
	AND password = $2
	`
	_ = sql
	return nil
}
