package db

import (
	"context"

	"github.com/ekowdd89/go-kit/internals/db/sqlc"
)



func (p *Postgres) FetchAll(ctx context.Context) ([]sqlc.Author, error) {
	return p.q.FetchAuthors(ctx)
}
