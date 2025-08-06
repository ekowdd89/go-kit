package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ekowdd89/go-kit/internals/db/sqlc"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)


type OptsFunc func(*Postgres) error


func WithConnectionString(dbUrl string) OptsFunc {
	return func(p *Postgres) (err error) {
		p.dbUrl = dbUrl
		p.db , err = otelsql.Open("postgres", dbUrl)
		return nil
	}
}

func WithTracer(name string) OptsFunc {
	return func(p *Postgres) (err error) {
		p.tracer = otel.Tracer(name)
		return
	}
}

type Postgres struct {
	dbUrl string
	db *sql.DB
	q *sqlc.Queries
	tracer trace.Tracer
	closers []func(context context.Context) error
}

func New(opts ...OptsFunc)(pg *Postgres, err error){
	pg = &Postgres{
		tracer: otel.Tracer("Postgres"),
	}
	for _, opt:= range opts {
		if err = opt(pg); err !=nil {
			return
		}
	}
	if pg.db == nil {
		return nil, fmt.Errorf("Missing connection database")
	}
	pg.q = sqlc.New(pg.db)
	return
}


func (p *Postgres) Close(ctx context.Context) (err error){
	for _, closer := range p.closers {
		err = errors.Join(err, closer(ctx))
	}
	return errors.Join(p.db.Close())
}