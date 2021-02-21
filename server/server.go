package server

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"local.packages/controller"
)

//Server is start server method
func Server() {
	engine := gin.Default()
	engine.Use(static.Serve("/", static.LocalFile("./static", false)))
	apiEngine := engine.Group("/api")
	{
		portfolioEngine := apiEngine.Group("/portfolio")
		{
			portfolioEngine.GET("", controller.GetPortfolio)
		}
	}
	engine.Run(":3000")
}
