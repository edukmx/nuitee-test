package httpx

import (
	"log/slog"
	"net/http"

	"github.com/edukmx/nuitee/config"
	"github.com/edukmx/nuitee/internal/ui/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config             *config.Config
	hotelsRatesHandler *handler.HotelsRatesHandler
}

func NewServer(
	config *config.Config,
	hotelsRatesHandler *handler.HotelsRatesHandler,
) *Server {

	return &Server{
		config:             config,
		hotelsRatesHandler: hotelsRatesHandler,
	}
}

func (s *Server) Run() {

	router := gin.Default()
	router.GET("/", s.Welcome)
	router.GET("/hotels", s.hotelsRatesHandler.List)

	err := router.Run(s.config.Port)

	if err != nil {
		slog.Error(err.Error())
	}
}

func (s *Server) Welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "welcome to api v1",
	})
}
