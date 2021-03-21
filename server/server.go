package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//Server is start server method
func Server() {
	createFireStoreJSON()
	if err := migration(); err != nil {
		panic("db connection error!")
	}
	engine := gin.Default()
	config := cors.DefaultConfig()
	runnningMode := os.Getenv("GIN_MODE")
	if runnningMode == "release" {
		fmt.Println("release mode")
		origins := strings.Split(os.Getenv("FRONTEND_DOMAINS"), ",")
		config.AllowOrigins = origins

	} else {
		fmt.Println("debug mode")
		config.AllowOrigins = []string{"*"}
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization"}
	engine.Use(cors.New(config))
	engine.Use(static.Serve("/", static.LocalFile("./static", false)))
	apiEngine := engine.Group("/api")
	{
		portfolioEngine := apiEngine.Group("/portfolios")
		{
			portfolioEngine.GET("/:id", getPortfolioHandler)
			portfolioEngine.GET("", getPortfolioListHandler)
			authGroup := portfolioEngine.Group("", authMiddleware())
			{
				authGroup.POST("", postPortfolioHandler)
				authGroup.PUT("/:id", putPortfolioHandler)
				authGroup.DELETE("/:id", deletePortfolioHandler)
			}
		}
	}
	engine.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	engine.Run(":" + port)
}
