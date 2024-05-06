package inventory

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/data"
	"github.com/steelthedev/irs/models"
)

type InventoryHandler struct {
	ProductService *models.ProductService
}

func (h InventoryHandler) RemoveStockQuantity(ctx *gin.Context) {
	var params data.QuanityDto
	if err := ctx.BindJSON(&params); err != nil {
		slog.Info("Invalid body request", "Error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Invalid body request",
			Code:    http.StatusBadRequest,
		})
		return
	}
	productID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := h.productService.RemoveQuantity(uint(productID), params.Units); err != nil {
		slog.Info("Could not Remove Quantity", "Error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "An error occured",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": "Success"})
}

func (h InventoryHandler) IncreaseStockQuantity(ctx *gin.Context) {
	var params data.QuanityDto
	if err := ctx.BindJSON(&params); err != nil {
		slog.Info("Invalid body request", "Error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Invalid body request",
			Code:    http.StatusBadRequest,
		})
		return
	}
	productID, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err := h.productService.AddQuantity(uint(productID), params.Units); err != nil {
		slog.Info("Could not Add Quantity", "Error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "An error occured",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": "Success"})
}
