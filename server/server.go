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
	engine := gin.Default()
	config := cors.DefaultConfig()
	runnningMode := os.Getenv("GIN_MODE")
	if runnningMode == "release" {
		fmt.Println("release mode")
	} else {
		fmt.Println("debug mode")
		config.AllowOrigins = []string{"http://localhost:5000"}
		engine.Use(cors.New(config))
	}
	engine.Use(static.Serve("/", static.LocalFile("./static", false)))
	apiEngine := engine.Group("/api")
	{
		portfolioEngine := apiEngine.Group("/portfolio")
		{
			portfolioEngine.GET("", getList)
		}
	}
	engine.Run(":3000")
}
