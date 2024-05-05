package products

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
)

type ProductsHandler struct {
	ProductService models.ProductService
}

func (h ProductsHandler) AddNewProduct(ctx *gin.Context) {
	var params models.AddProductParams

	if err := ctx.BindJSON(&params); err != nil {
		slog.Info("Invalid Body Request", "error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Invalid Body Request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	product, err := h.ProductService.AddNewProduct(params)
	if err != nil {
		slog.Info("An error Occured", "error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.IndentedJSON(201, &product)
}
