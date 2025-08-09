package db

import (
	"context"
	"database/sql"
	"github.com/ekowdd89/go-kit/internals/db/sqlc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)



func (p *Postgres) FetchAll(ctx context.Context) ([]sqlc.Author, error) {
	_, span := p.tracer.Start(ctx, "FetchAll",
		trace.WithAttributes(attribute.String("tracer", "Postgres")),
	)
	span.SetAttributes(
		attribute.String("trace_id", "FetchAll"),
		attribute.String("span_id", "FetchAllSpanId"),
	)
	defer span.End()

	return p.q.FetchAuthors(ctx)
}

func (p *Postgres) FetchById(ctx context.Context, id int64) (sqlc.Author, error) {
	_, span := p.tracer.Start(ctx, "FetchById",
		trace.WithAttributes(attribute.String("tracer", "Postgres")),
	)
	span.SetAttributes(
		attribute.String("trace_id", "FetchById"),
		attribute.String("span_id", "FetchByIdSpanId"),
	)
	defer span.End()

	return p.q.FetchAuthor(ctx, int32(id))
}

func (p *Postgres) Create(ctx context.Context, name sql.NullString ) (sqlc.Author, error) {
	_, span := p.tracer.Start(ctx, "Create",
		trace.WithAttributes(attribute.String("tracer", "Postgres")),
	)
	span.SetAttributes(
		attribute.String("trace_id", "Create"),
		attribute.String("span_id", "CreateSpanId"),
	)
	defer span.End()

	return p.q.CreateAuthor(ctx, name)
}

func (p *Postgres) Update(ctx context.Context, request sqlc.UpdateAuthorParams) (sqlc.Author, error) {
	_, span := p.tracer.Start(ctx, "Update",
		trace.WithAttributes(attribute.String("tracer", "Postgres")),
	)
	span.SetAttributes(
		attribute.String("trace_id", "Update"),
		attribute.String("span_id", "UpdateSpanId"),
	)
	defer span.End()

	return p.q.UpdateAuthor(ctx, request)
}

func (p *Postgres) Delete(ctx context.Context, id int64) (sqlc.Author, error) {
	_, span := p.tracer.Start(ctx, "Delete",
		trace.WithAttributes(attribute.String("tracer", "Postgres")),
	)
	span.SetAttributes(
		attribute.String("trace_id", "Delete"),
		attribute.String("span_id", "DeleteSpanId"),
	)
	defer span.End()

	return p.q.DeleteAuthor(ctx, int32(id))
}