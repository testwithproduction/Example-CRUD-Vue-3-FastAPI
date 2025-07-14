package routes

import (
	"api-go/handlers"
	"api-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORS())

	// Serve static files
	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*.html")

	// Root route to serve index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// API routes
	api := r.Group("/api")
	{
		// Product routes
		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.POST("", handlers.CreateProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
		}
	}

	return r
} 