package controller

import (
	"github.com/gin-gonic/gin"
	"local.packages/service"
)

//GetPortfolioList is return Portfoliolist
func GetPortfolioList(c *gin.Context) {
	portfolioList := service.GetPortfolioList()
	c.JSON(200, portfolioList)
}
