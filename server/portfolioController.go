package server

import (
	"github.com/gin-gonic/gin"
)

func getList(c *gin.Context) {
	portfolioList := getPortfolioList()
	c.JSON(200, portfolioList)
}
