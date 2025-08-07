package otel

import (
	"github.com/ekowdd89/go-kit/internals/db/author"
	"go.opentelemetry.io/otel"
)


var Tracer = otel.Tracer("otelWrapper")

//go:generate otelwrap --out author-contract.go . author.AuthorContract
var _ author.AuthorContract