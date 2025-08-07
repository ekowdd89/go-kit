package author

import (
	// "io"
	"context"

	"github.com/ekowdd89/go-kit/internals/db/sqlc"
)



type AuthorContract interface {
	// io.Closer
	FetchAll(ct context.Context)([]sqlc.Author, error)
	// FetchById(int64) (Author, error)
	// Create(AuthorRequest) (Author, error)
	// Update(AuthorUpdateRequest) (Author, error)
	// Delete(AuthorDeleteRequest) (Author, error)
}

type Author struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type AuthorRequest struct {
	Name string
}
type AuthorUpdateRequest struct {
	Id int64
	Name string
}

type AuthorDeleteRequest struct {
	Id int64
}