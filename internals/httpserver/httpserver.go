package httpserver

import "github.com/gin-gonic/gin"




type OptsFunc func(*HTTPServer) error


type HTTPServer struct {
	handler gin.Engine
}

func (h HTTPServer) HealthCheck() (hdl HTTPServer, err error) {
	return
}