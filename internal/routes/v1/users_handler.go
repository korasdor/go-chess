package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-chess/internal/domain"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("users", h.userIdentity)
	{
		users.GET("/current", h.getCurrentUser)
		users.GET("/:id", h.getUser)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)

		wallet := users.Group(":id/wallet")
		{
			wallet.POST("/", h.addWallet)
			wallet.GET("/", h.getAllWallets)
			wallet.GET("/:walletId", h.getWallet)
			wallet.DELETE("/", h.deleteWallet)

			balance := wallet.Group("balance")
			{
				balance.GET("/:addr", h.getBalance)
			}
		}
	}
}

// get user...
func (h *Handler) getCurrentUser(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.UsersService.GetUser(userId)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// get user...
func (h *Handler) getUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.UsersService.GetUser(userId)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// update user...
func (h *Handler) updateUser(ctx *gin.Context) {
	var user domain.UserData

	if err := ctx.BindJSON(&user); err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user.Id = userId

	user, err = h.services.UsersService.UpdateUser(user)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// delete user...
func (h *Handler) deleteUser(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

// add wallet...
func (h *Handler) addWallet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

// get all wallets...
func (h *Handler) getAllWallets(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

// get wallet...
func (h *Handler) getWallet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

// delete wallet...
func (h *Handler) deleteWallet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

// get balance...
func (h *Handler) getBalance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}
