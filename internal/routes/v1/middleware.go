package v1

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	userId, err := h.parseAuthHeader(ctx)
	if err != nil {
		newResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func (h *Handler) parseAuthHeader(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.ParseJWT(headerParts[1])
}

func getUserId(c *gin.Context) (int, error) {
	id, err := getIdByContext(c, userCtx)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(id)
}

func getIdByContext(c *gin.Context, context string) (string, error) {
	var idStr string
	idFromCtx, ok := c.Get(context)
	if !ok {
		return idStr, errors.New("userCtx not found")
	}

	idStr, ok = idFromCtx.(string)
	if !ok {
		return idStr, errors.New("userCtx is of invalid type")
	}

	return idStr, nil
}
