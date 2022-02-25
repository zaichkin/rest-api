package users

import "context"

type Storage interface {
	CreateUser(ctx context.Context, dto SingUpDTO) error
	GetUser(ctx context.Context, dto SingInDTO) error
}
