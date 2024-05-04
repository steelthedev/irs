package accounts

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/utils"
	"gorm.io/gorm"
)

type AccountHandler struct {
	DB *gorm.DB
}

func (h AccountHandler) GetUserProfile(ctx *gin.Context) {
	// Get userId from token
	userId, err := utils.ExtractIdFromToken(ctx)
	if err != nil {
		slog.Info(err.Error())
		ctx.Error(&data.AppHttpErr{Message: "An error occured", Code: http.StatusInternalServerError})
		return
	}

	// Fetch user from database
	user, err := utils.GetUserById(userId, h.DB)
	if err != nil {
		ctx.Error(&data.AppHttpErr{Message: "An error occured", Code: http.StatusInternalServerError})
		return
	}

	ctx.IndentedJSON(200, &user)
}