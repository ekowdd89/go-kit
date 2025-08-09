package todo

import (
	"context"

	"github.com/ekowdd89/go-kit/internals/db/sqlc"
)


type TodoContract interface {
	GetByID(ctx context.Context, id int64) (sqlc.Todo, error)
	GetAll(ctx context.Context) ([]sqlc.Todo, error)
	Create(ctx context.Context, todo sqlc.CreateTodoParams) (sqlc.Todo, error)
	Update(ctx context.Context, todo sqlc.UpdateTodoParams) (sqlc.Todo, error)
	Delete(ctx context.Context, id int64) (sqlc.Todo, error)
}


