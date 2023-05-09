package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-chess/internal/config"
	v1 "github.com/korasdor/go-chess/internal/routes/v1"
	"github.com/korasdor/go-chess/internal/services"
	"github.com/korasdor/go-commons/auth"
	"github.com/korasdor/go-commons/limiter"
)

type Handler struct {
	services     *services.Services
	config       *config.Config
	tokenManager auth.TokenManager
}

func NewHandler(services *services.Services, config *config.Config, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		config:       config,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(h.config.Gin.GinMode)

	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		limiter.Limit(h.config.Limiter.RPS, h.config.Limiter.Burst, h.config.Limiter.TTL),
		corsMiddleware,
	)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	h.initApp(router)

	return router
}

func (h *Handler) initApp(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.config, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
