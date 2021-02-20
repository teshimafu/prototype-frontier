package server

import "github.com/gin-gonic/gin"

//Server is start server method
func Server() {
	engine := gin.Default()
	engine.Static("/", "./static")
	engine.Run(":3000")
}
