package pg

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, connStr string) (*Postgres, error) {
	var initErr error
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connStr)
		if err != nil {
			initErr = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &Postgres{Db: db}
	})

	if initErr != nil {
		return nil, initErr
	}

	return pgInstance, nil
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
