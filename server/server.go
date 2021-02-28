package server

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//Server is start server method
func Server() {
	if err := migration(); err != nil {
		panic("db connection error!")
	}
	engine := gin.Default()
	config := cors.DefaultConfig()
	runnningMode := os.Getenv("GIN_MODE")
	if runnningMode == "release" {
		fmt.Println("release mode")
	} else {
		fmt.Println("debug mode")
		config.AllowOrigins = []string{"*"}
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	engine.Use(cors.New(config))
	engine.Use(static.Serve("/", static.LocalFile("./static", false)))
	apiEngine := engine.Group("/api")
	{
		portfolioEngine := apiEngine.Group("/portfolio")
		{
			portfolioEngine.GET("/:id", getPortfolioHandler)
			portfolioEngine.GET("", getPortfolioListHandler)
			portfolioEngine.POST("", postPortfolioHandler)
		}
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	engine.Run(":" + port)
}
