package main

import (
	"context"
	"fmt"

	"github.com/ekowdd89/go-kit/internals/db"
	"github.com/ekowdd89/go-kit/internals/otel"
	"github.com/gin-gonic/gin"
	"os"
)



func main(){
	g:= gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		otel.InitTracer(context.Background())
		_, err:= db.New(
			db.WithConnectionString("postgres://testing:testing@postgresql:5434/testing?sslmode=disable"),
			db.WithTracer("main postgres"),
		)
		if err !=nil {
			fmt.Println(err.Error())
		}
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err:= g.Run(fmt.Sprintf(":%s", os.Getenv("APP_LISTEN"))); err != nil {
		panic(err)
	}
}