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
	// postgres://testing:testing@postgresql:5434/testing?sslmode=disable
	// postgres://postgres:cempakamasp23@127.0.0.1:5434/sertif_tes?sslmode=disable
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
	fmt.Println("server started listening on port 8080")
	// g:= gin.Default()
	// g.GET("/", func(ctx *gin.Context) {
	// 	otel.InitTracer(context.Background())
	// 	_, err:= db.New(
	// 		db.WithConnectionString("postgres://testing:testing@postgresql:5434/testing?sslmode=disable"),
	// 		db.WithTracer("main postgres"),
	// 	)
	// 	if err !=nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// if err:= g.Run(fmt.Sprintf(":%s", os.Getenv("APP_LISTEN"))); err != nil {
	// 	panic(err)
	// }
}