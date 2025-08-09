package main

import (
	"net"

	"context"
	"fmt"
	"os"

	"github.com/ekowdd89/go-kit/internals/db"
	"github.com/ekowdd89/go-kit/pkg/httpserver"
	otelwrap "github.com/ekowdd89/go-kit/internals/otel"
)

func main(){
	l, err:= net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_LISTEN")))
	if err != nil {
		panic(err)
	}
	p,err:= db.New(
		db.WithConnectionString(os.Getenv("DB_URL")),
		db.WithTracer("main postgres"),
	)
	if err !=nil {
		println(err.Error())
	}
	h, err:= httpserver.New(
		httpserver.WithListener(l),
		httpserver.WithAuthorContract(
			otelwrap.NewAuthorContractWrapper(p, otelwrap.Tracer, "author"),
		),
		httpserver.WithTracer("main httpserver"),
	)
	if err !=nil {
		h.Close(context.Background())
	}
	if err := h.Start(context.Background()); err != nil {
		h.Close(context.Background())
	}
	fmt.Println("server started listening on port ", os.Getenv("APP_LISTEN"))
}