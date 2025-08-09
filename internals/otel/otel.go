package otel

import (
	"github.com/ekowdd89/go-kit/internals/db/author"
	"github.com/ekowdd89/go-kit/internals/db/todo"
	"go.opentelemetry.io/otel"
)


var Tracer = otel.Tracer("otelWrapper")

//go:generate otelwrap --out author-contract.go . author.AuthorContract
var _ author.AuthorContract

//go:generate otelwrap --out todo-contract.go . todo.TodoContract
var _ todo.TodoContract