package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPortfolioListHandler(c *gin.Context) {
	portfolioList, err := getPortfolioList()
	if err != nil {
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.JSON(200, portfolioList)
}

func postPortfolioHandler(c *gin.Context) {
	var p Portfolio
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := postPortfolio(&p); err != nil {
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.Status(201)
}
