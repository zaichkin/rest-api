package postgresql

import (
	"context"
	"fmt"
	"market/internal/config"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Connect interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewConnection(countRepeat int, conf *config.Config) (*pgxpool.Pool, error) {
	var db *pgxpool.Pool
	var err error

	param := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.PgUser, conf.PgPasswd, conf.PgHost, conf.PgPort, conf.PgDB)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < countRepeat; i++ {

		fmt.Print("Connection attempt: ")
		db, err = pgxpool.Connect(ctx, param)

		if err != nil {
			time.Sleep(time.Second)
			fmt.Println("Fail!")
			continue
		}

		fmt.Println("Success!")
		return db, nil
	}

	return nil, err
}
