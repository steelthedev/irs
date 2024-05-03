package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/connections"
	"github.com/steelthedev/irs/middlewares"
)

func main() {

	app := gin.Default()

	_, err := connections.InitDb()
	if err != nil {
		slog.Info(err.Error())
	}

	app.Use(middlewares.ErrorHandler())

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{
			"version": "v1",
			"app":     "IRS",
		})
	})

	app.Run(":3000")

}
