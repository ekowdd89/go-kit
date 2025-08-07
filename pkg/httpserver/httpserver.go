package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"

	// "encoding/json"

	// "github.com/ekowdd89/go-kit/internals/db"
	"github.com/ekowdd89/go-kit/internals/db/author"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)




type OptsFunc func(*HTTPServer) error


func WithListener(l net.Listener) OptsFunc {
	return func(h *HTTPServer) (err error) {
		h.listener = l
		return
	}
}


type HTTPServer struct {
	authorRepo author.AuthorContract

	engine *gin.Engine
	tracerName string
	server *http.Server
	listener net.Listener
}


func WithAuthorContract(repo author.AuthorContract) OptsFunc {
	return func(p *HTTPServer) (err error) {
		p.authorRepo = repo
		return
	}
}



func New(opts ...OptsFunc) (h *HTTPServer, err error) {
	h =&HTTPServer{
		engine: gin.Default(),
		tracerName: "HTTPServer",

	}
	for _, opt:= range opts {
		if err = opt(h); err !=nil {
			return
		}
	}
	err = h.buildServer()
	return
}

func (h *HTTPServer) buildServer() (err error) {
	// h.engine.
	h.engine.Use(gin.Logger())
	h.engine.Use(gin.Recovery())
	h.engine.Use(cors.Default())
	h.Ingpoh().HealthCheck()

	h.server = &http.Server{
		Addr: ":8080",
		Handler: h.engine,
	}
	return
}

func (h *HTTPServer) Ingpoh() *HTTPServer {
	h.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Api": "test api",
			"Version": "1.0.0",
		})
	})
	return h
}

func (h *HTTPServer) HealthCheck() *HTTPServer {
	h.engine.GET("/health", func(ctx *gin.Context) {
		d, err :=h.authorRepo.FetchAll(ctx.Context())
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		authors:= []author.Author{}
		for _, v := range d {
			authors = append(authors, author.Author{
				Id: int64(v.ID),
				Name: v.Name.String,
			})
		}
		println("d", authors)
		ctx.JSON(200, gin.H{
			"message": "ok",
			"data": authors,
		})
	})
	return h
}

func (h HTTPServer) Start(ctx context.Context)(err error){
	go func() {
		<-ctx.Done()
		err = errors.Join(err, h.server.Shutdown(ctx))
	}()
	if h.listener == nil {
		return h.server.ListenAndServe()
	}
	return errors.Join(err, h.server.Serve(h.listener))
}

func (h HTTPServer) Close(ctx context.Context) (err error) {
	return h.server.Shutdown(ctx)
}
