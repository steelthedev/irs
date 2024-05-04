package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
	"github.com/steelthedev/irs/utils"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()
			var appErr *data.AppHttpErr
			switch err.Err.(type) {
			case *data.AppHttpErr:
				appErr = err.Err.(*data.AppHttpErr)
			default:
				appErr = &data.AppHttpErr{
					Message: err.Error(),
					Code:    http.StatusInternalServerError,
				}
			}
			ctx.IndentedJSON(appErr.Code, appErr)
		}
	}
}

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := utils.ExtractToken(ctx)
		if token == "" {
			slog.Info("Authentication credentials are missing")
			ctx.AbortWithError(http.StatusUnauthorized, &data.AppHttpErr{
				Message: "Authentication credentials are missing",
				Code:    http.StatusUnauthorized,
			})
			return
		}

		if err := utils.TokenValid(ctx); err != nil {
			slog.Info("Invalid token", "Error", err.Error())
			ctx.AbortWithError(http.StatusUnauthorized, &data.AppHttpErr{
				Message: "Invalid Token",
				Code:    http.StatusUnauthorized,
			})
			return
		}

		ctx.Next()
	}
}

func IsAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.ExtractIdFromToken(ctx)
		if err != nil {
			slog.Info("Could not get user details", "Error", err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, &data.AppHttpErr{
				Message: "An error occured",
				Code:    http.StatusInternalServerError,
			})
			return
		}

		// Get user from Id
		user, err := models.GetUserById(userId, db)
		if err != nil {
			slog.Info("User not recognixed", "Error", err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, &data.AppHttpErr{
				Message: "Logged in User Not recognized",
				Code:    http.StatusInternalServerError,
			})
			return
		}

		// Check if user is admin

		if user.Role != models.AdminRole {
			slog.Info("Not Authorized to access resource")
			ctx.AbortWithError(http.StatusUnauthorized, &data.AppHttpErr{
				Message: "Not Authorized to acess this resources",
				Code:    http.StatusUnauthorized,
			})
			return
		}
		ctx.Next()
	}
}
