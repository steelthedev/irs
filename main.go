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

	// app.Use(.ErrorHandler())

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{
			"version": "v1",
			"app":     "IRS",
		})
	})

	// Services

	userServices := &models.UserService{
		DB: db,
	}

	productService := &models.ProductService{
		DB: db,
	}

	// Handlers Structs
	accountHandler := &accounts.AccountHandler{
		UserService: userServices,
	}

	// Auth App Handler
	authHandler := &auth.AuthHandler{
		UserServices: *userServices,
	}

	adminHandler := &admin.AdminHandler{
		DB: db,
	}

	productHandler := &products.ProductsHandler{
		ProductService: *productService,
	}

	// Middleware Hanlder
	middleWareHandler := middlewares.MiddlewareHandler{
		UserServices: userServices,
	}
	//Routes

	// Auth Routes
	authRoutes := app.Group("auth")
	authRoutes.POST("/register", authHandler.CreateUser)
	authRoutes.POST("/login", authHandler.Login)

	// Admin Routes
	adminRoutes := app.Group("admin")
	adminRoutes.Use(middleWareHandler.IsAuthenticated())
	adminRoutes.Use(middleWareHandler.IsAdmin())
	adminRoutes.GET("/users", adminHandler.GetAllUsers)

	// Account Routes
	accountRoutes := app.Group("account", middleWareHandler.IsAuthenticated())
	accountRoutes.GET("/profile", accountHandler.GetUserProfile)

	// Product Routes
	productRoutes := app.Group("product", middleWareHandler.IsAuthenticated())
	productRoutes.POST("/add", productHandler.AddNewProduct)
	productRoutes.DELETE("/delete/:id", productHandler.DeleteProduct)
	productRoutes.GET("/", productHandler.GetAllProducts)

	// Start app
	app.Run(":3000")

}
