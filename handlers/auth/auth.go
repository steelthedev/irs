package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
	"github.com/steelthedev/irs/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h AuthHandler) CreateUser(ctx *gin.Context) {
	var params data.RegisterUser

	if err := ctx.BindJSON(&params); err != nil {
		slog.Warn(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Invalid Body Request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := params.Validate(); err != nil {
		slog.Info(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	hashedPwd, err := utils.HashPassword(params.Password)
	if err != nil {
		slog.Info(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "An error ocurred while hashing password",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	user := &models.User{
		Email:     params.Email,
		Password:  string(hashedPwd),
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	if result := h.DB.Create(&user); result.Error != nil {
		slog.Info(result.Error.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "User could not be created",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	ctx.IndentedJSON(200, gin.H{
		"message": "User created successfully",
	})

}
