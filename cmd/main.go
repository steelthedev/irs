package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/connections"
	"github.com/steelthedev/irs/handlers/auth"
	"github.com/steelthedev/irs/middlewares"
)

func main() {

	app := gin.Default()

	db, err := connections.InitDb()
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

	// Auth App Handler
	authHandler := &auth.AuthHandler{
		DB: db,
	}

	//Routes

	// Auth Routes
	authRoutes := app.Group("auth")
	authRoutes.POST("/register", authHandler.CreateUser)

	// Start app
	app.Run(":3000")

}
