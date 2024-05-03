package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
	"github.com/steelthedev/irs/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h AuthHandler) CreateUser(ctx *gin.Context) {
	var params data.RegisterUser

	// Bind Body
	if err := ctx.BindJSON(&params); err != nil {
		slog.Warn(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Invalid Body Request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Validate Inputed parameters
	if err := params.Validate(); err != nil {
		slog.Info(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Check if user already exists
	if utils.CheckUserExistsWithEmail(params.Email, h.DB) {
		slog.Info("User with email already exists", "Email", params.Email)
		ctx.Error(&data.AppHttpErr{
			Message: "User with this email already exists",
			Code:    500,
		})
		return
	}

	// Hash inputed password with bCrypt
	hashedPwd, err := utils.HashPassword(params.Password)
	if err != nil {
		slog.Info(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "An error ocurred while hashing password",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Create User Object
	user := &models.User{
		Email:     params.Email,
		Password:  string(hashedPwd),
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	// Add user to database
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

func (h AuthHandler) Login(ctx *gin.Context) {

	var params data.LoginUser

	if err := ctx.BindJSON(&params); err != nil {
		slog.Warn(err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Invalid Body Request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Check if user with email exists

	if !utils.CheckUserExistsWithEmail(params.Email, h.DB) {
		slog.Info("User with email does not exist", "Email", params.Email)
		ctx.Error(&data.AppHttpErr{
			Message: "User Does Not Exist",
			Code:    http.StatusNotFound,
		})
		return
	}

	// Get user from DB
	var user models.User

	if result := h.DB.Where("Email=?", params.Email).First(&user); result.Error != nil {
		slog.Info(result.Error.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "An unexpected error occured",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Compare passwords

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		ctx.Error(&data.AppHttpErr{
			Message: "Incorrect Password.",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user)
	if err != nil {
		ctx.Error(&data.AppHttpErr{
			Message: "An error occured",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Send token to user
	ctx.IndentedJSON(200, gin.H{
		"access_token": token,
	})
}
