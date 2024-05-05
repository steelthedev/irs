package main

import (
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steelthedev/irs/connections"
	"github.com/steelthedev/irs/handlers/accounts"
	"github.com/steelthedev/irs/handlers/admin"
	"github.com/steelthedev/irs/handlers/auth"
	"github.com/steelthedev/irs/handlers/products"
	"github.com/steelthedev/irs/middlewares"
	"github.com/steelthedev/irs/models"
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
	if err := godotenv.Load(".env"); err != nil {
		slog.Info(err.Error())
	}
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

	accountHandler := &accounts.AccountHandler{
		DB: db,
	}

	// Auth App Handler
	authHandler := &auth.AuthHandler{
		DB: db,
	}

	adminHandler := &admin.AdminHandler{
		DB: db,
	}

	productService := &models.ProductService{
		DB: db,
	}
	productHandler := &products.ProductsHandler{
		ProductService: *productService,
	}
	//Routes

	// Auth Routes
	authRoutes := app.Group("auth")
	authRoutes.POST("/register", authHandler.CreateUser)
	authRoutes.POST("/login", authHandler.Login)

	// Admin Routes
	adminRoutes := app.Group("admin")
	adminRoutes.Use(middlewares.IsAuthenticated())
	adminRoutes.Use(middlewares.IsAdmin(db))
	adminRoutes.GET("/users", adminHandler.GetAllUsers)

	// Account Routes
	accountRoutes := app.Group("account", middlewares.IsAuthenticated())
	accountRoutes.GET("/profile", accountHandler.GetUserProfile)

	// Product Routes
	productRoutes := app.Group("product", middlewares.IsAuthenticated())
	productRoutes.POST("/add", productHandler.AddNewProduct)

	// Start app
	app.Run(":3000")

}
