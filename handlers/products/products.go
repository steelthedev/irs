package products

import (
	"log/slog"
	"net/http"
	"strconv"

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
			Message: "An error occured",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.IndentedJSON(201, &product)
}

func (h ProductsHandler) DeleteProduct(ctx *gin.Context) {
	// Get Product ID
	productID, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		slog.Info("error parsing id to uint", "error", err.Error())
	}

	// Delete Logic
	if err := h.ProductService.DeleteProduct(uint(productID)); err != nil {
		slog.Info("An error occured", "Error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "An error occured",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Product Deleted succesfuuly"})
}

func (h ProductsHandler) GetAllProducts(ctx *gin.Context) {

	products, err := h.ProductService.GetAllProduct()
	if err != nil {
		slog.Info("Could not fetch products", "error", err.Error())
		ctx.Error(&data.AppHttpErr{
			Message: "Products coudl not be fetched",
			Code:    http.StatusNotFound,
		})
		return
	}
	ctx.IndentedJSON(200, &products)
}
