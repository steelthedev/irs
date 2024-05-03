package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
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
