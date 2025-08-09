package httpserver

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// "encoding/json"

	// "github.com/ekowdd89/go-kit/internals/db"
	"github.com/ekowdd89/go-kit/internals/db/author"
	"github.com/ekowdd89/go-kit/internals/db/sqlc"
	"github.com/ekowdd89/go-kit/internals/tracing"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/ekowdd89/go-kit/internals/tracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/ekowdd89/go-kit/pkg/httpserver/middleware"
)




type OptsFunc func(*HTTPServer) error


func WithListener(l net.Listener) OptsFunc {
	return func(h *HTTPServer) (err error) {
		h.listener = l
		return
	}
}

func WithTracer(tracerName string) OptsFunc {
	return func(h *HTTPServer) (err error) {
		h.tracerName = tracerName
		h.tracer = otel.Tracer(tracerName)
		return
	}
}
type HTTPServer struct {
	authorRepo author.AuthorContract

	engine *gin.Engine
	tracerName string

	tracer trace.Tracer

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
	tp, err := tracing.New(context.Background())
	if err != nil {
		return
	}
	// defer func() { _ = tp.Shutdown(context.Background()) }()
	h.tracer = tp.Tracer("HTTPServer")
	err = h.buildServer()
	return
}

func (h *HTTPServer) buildServer() (err error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	h.engine.Use(otelgin.Middleware(h.tracerName))
	h.engine.Use(middleware.ZeroNew())
	h.engine.Use(gin.Logger())
	h.engine.Use(gin.Recovery())
	h.engine.Use(cors.Default())
	h.Ingpoh().HealthCheck().FetchAuthorById().CreateAuthor().UpdateAuthor().DeleteAuthor()

	h.server = &http.Server{
		Handler: h.engine,
	}
	return
}

func (h *HTTPServer) Ingpoh() *HTTPServer {
	h.engine.GET("/", func(ctx *gin.Context) {
		_, span := h.tracer.Start(ctx, "HTTP GET /Ingpoh")
		defer span.End()
		println("span", span)
		ctx.JSON(200, gin.H{
			"Api": "test api",
			"Version": "1.0.0",
		})
	})
	return h
}

func (h *HTTPServer) HealthCheck() *HTTPServer {
	h.engine.GET("/health", func(ctx *gin.Context) {
		_ctx, span := h.tracer.Start(ctx, "HTTP GET /health")
		println("span", span)
		defer span.End()
		d, err :=h.authorRepo.FetchAll(_ctx)
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

func (h *HTTPServer) FetchAuthorById() *HTTPServer {
	h.engine.GET("/author/:id", func(ctx *gin.Context) {
		_, span:= h.tracer.Start(ctx, "HTTP GET /FetchAuthorById")
		defer span.End()
		_id:= ctx.Param("id")
		id, err := strconv.Atoi(_id)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		d, err :=h.authorRepo.FetchById(ctx, int64(id))
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		author:= author.Author{
			Id: int64(d.ID),
			Name: d.Name.String,
		}
		ctx.JSON(200, gin.H{
			"message": "ok",
			"data": author,
		})
	})
	return h
}
func (h *HTTPServer) CreateAuthor() *HTTPServer {
	h.engine.POST("/author", func(ctx *gin.Context) {
		_, span:= h.tracer.Start(ctx, "HTTP POST /CreateAuthor")
		defer span.End()
		var req author.AuthorRequest
		var name sql.NullString
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		name = sql.NullString{
			String: req.Name,
			Valid: true,
		}
		println("name", name.String)
		d, err :=h.authorRepo.Create(ctx, name)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "ok",
			"data": d,
		})
	})
	return h
}

func (h *HTTPServer) UpdateAuthor() *HTTPServer {
	h.engine.PUT("/author/:id", func(ctx *gin.Context) {
		_, span:= h.tracer.Start(ctx, "HTTP PUT /UpdateAuthor")
		defer span.End()
		_id:= ctx.Param("id")
		id, err := strconv.Atoi(_id)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		var req author.AuthorUpdateRequest
		var name sql.NullString
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		name = sql.NullString{
			String: req.Name,
			Valid: true,
		}
		d, err :=h.authorRepo.Update(ctx,
			sqlc.UpdateAuthorParams{
				ID: int32(id),
				Name: name,
			},
		)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "ok",
			"data": d,
		})
	})
	return h
}

func (h *HTTPServer) DeleteAuthor() *HTTPServer {
	h.engine.DELETE("/author/:id", func(ctx *gin.Context) {
		_, span:= h.tracer.Start(ctx, "HTTP DELETE /DeleteAuthor")
		defer span.End()
		_id:= ctx.Param("id")
		id, err := strconv.Atoi(_id)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		d, err :=h.authorRepo.Delete(ctx, int64(id))
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "ok",
			"data": d,
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
