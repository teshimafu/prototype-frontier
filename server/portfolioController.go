package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func getPortfolioHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "parse param error"})
		return
	}
	portfolio, err := getPortfolio(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"message": "データが見つかりませんでした"})
			return
		}
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.JSON(200, portfolio)
}

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

	var portfolio, err = postPortfolio(&p)
	if err != nil {
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.JSON(201, portfolio)
}

func putPortfolioHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "parse param error"})
		return
	}
	var p Portfolio
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	portfolio, err := putPortfolio(id, &p)
	if err != nil {
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.JSON(200, portfolio)
}

func deletePortfolioHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "parse param error"})
		return
	}
	err = deletePortfolio(id)
	if err != nil {
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.Status(204)
}
