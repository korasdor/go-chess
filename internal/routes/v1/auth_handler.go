package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-chess/internal/domain"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("auth")
	{
		auth.POST("user/sign-up", h.userSignUp)
		auth.POST("user/sign-in", h.userSignIn)
		auth.POST("user/refresh", h.userRefreshToken)
	}
}

func (h *Handler) userSignUp(ctx *gin.Context) {
	var signUpData domain.SignUpData

	if err := ctx.BindJSON(&signUpData); err != nil {
		newResponse(ctx, http.StatusBadRequest, domain.ErrBadRequestFormat.Error())
		return
	}

	err := h.services.AuthorizationService.SignUp(signUpData)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

func (h *Handler) userSignIn(ctx *gin.Context) {
	var signInData domain.SignInData

	if err := ctx.BindJSON(&signInData); err != nil {
		newResponse(ctx, http.StatusBadRequest, domain.ErrBadRequestFormat.Error())
		return
	}

	token, err := h.services.AuthorizationService.SignIn(signInData)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.Auth.JWT.RefreshTokenTTL.Seconds()),
		HttpOnly: true,
	}

	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(http.StatusOK, token)
}

func (h *Handler) userRefreshToken(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("refreshToken")
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.AuthorizationService.RefreshTokens(cookie.Value)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	cookie = &http.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
		MaxAge:   int(h.config.Auth.JWT.RefreshTokenTTL.Seconds()),
		HttpOnly: true,
	}

	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(http.StatusOK, token)
}
