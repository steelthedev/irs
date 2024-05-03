package admin

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
	"gorm.io/gorm"
)

type AdminHandler struct {
	DB *gorm.DB
}

func (h AdminHandler) GetAllUsers(ctx *gin.Context) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		slog.Info(result.Error.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Could not load data",
			Code:    500,
		})
		return
	}

	ctx.IndentedJSON(200, &users)
}
