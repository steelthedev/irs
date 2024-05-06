package accounts

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
	"github.com/steelthedev/irs/tokens"
)

type AccountHandler struct {
	UserService *models.UserService
}

func (h AccountHandler) GetUserProfile(ctx *gin.Context) {
	// Get userId from token
	userId, err := tokens.ExtractIdFromToken(ctx)
	if err != nil {
		slog.Info(err.Error())
		ctx.Error(&data.AppHttpErr{Message: "An error occured", Code: http.StatusInternalServerError})
		return
	}

	// Fetch user from database
	user, err := h.UserService.GetUserById(userId)
	if err != nil {
		ctx.Error(&data.AppHttpErr{Message: "An error occured", Code: http.StatusInternalServerError})
		return
	}

	ctx.IndentedJSON(200, &user)
}
