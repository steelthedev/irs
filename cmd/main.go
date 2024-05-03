package main

import (
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/irs/connections"
	"github.com/steelthedev/irs/handlers/admin"
	"github.com/steelthedev/irs/handlers/auth"
	"github.com/steelthedev/irs/middlewares"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
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

	adminHandler := &admin.AdminHandler{
		DB: db,
	}

	//Routes

	// Auth Routes
	authRoutes := app.Group("auth")
	authRoutes.POST("/register", authHandler.CreateUser)

	// Admin Routes
	adminRoutes := app.Group("admin")
	adminRoutes.GET("/users", adminHandler.GetAllUsers)

	// Start app
	app.Run(":3000")

}
