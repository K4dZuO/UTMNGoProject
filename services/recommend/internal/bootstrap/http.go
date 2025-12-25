package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"

	api "recommend_service/internal/api/http"
)

type HTTPServer struct {
	addr   string
	engine *gin.Engine
}

func initHTTP(addr string, service api.RecommendService) *HTTPServer {
	router := gin.Default()

	handler := api.New(service)

	router.GET("/top", handler.GetTop)

	return &HTTPServer{
		addr:   addr,
		engine: router,
	}
}

func (s *HTTPServer) Run() error {
	log.Printf("HTTP server listening on %s", s.addr)
	return s.engine.Run(s.addr)
}
