package author

import (
	// "io"
	"context"
	"database/sql"

	"github.com/ekowdd89/go-kit/internals/db/sqlc"
)



type AuthorContract interface {
	// io.Closer
	FetchAll(ct context.Context)([]sqlc.Author, error)
	FetchById(ctx context.Context, id int64) (sqlc.Author, error)
	Create(ctx context.Context, name sql.NullString) (sqlc.Author, error)
	Update(ctx context.Context, request sqlc.UpdateAuthorParams) (sqlc.Author, error)
	Delete(ctx context.Context, id int64) (sqlc.Author, error)
}

type Author struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type AuthorRequest struct {
	Name string `json:"name"`
}
type AuthorUpdateRequest struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type AuthorDeleteRequest struct {
	Id int64
}