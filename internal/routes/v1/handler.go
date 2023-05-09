package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-chess/internal/config"
	"github.com/korasdor/go-chess/internal/services"
	"github.com/korasdor/go-commons/auth"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initUsersRoutes(v1)
	}
}
