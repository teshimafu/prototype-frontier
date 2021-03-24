package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr"
)

func getPortfolioHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "parse param error"})
		return
	}
	portfolio, err := getPortfolio(id)
	if err != nil {
		if err.Error() == dbr.ErrNotFound.Error() {
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
	uid, isExist := c.Get("AuthorizedUID")
	if isExist == false {
		c.JSON(500, gin.H{"message": "UID取得失敗"})
		return
	}
	p.UID = uid.(string)
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

	portfolio, err := getPortfolio(id)
	if err != nil {
		if err.Error() == dbr.ErrNotFound.Error() {
			c.JSON(404, gin.H{"message": "データが見つかりませんでした"})
			return
		}
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}

	if uid, isExist := c.Get("AuthorizedUID"); isExist == false {
		c.JSON(500, gin.H{"message": "UID取得失敗"})
		return
	} else if p.UID != "" && uid != p.UID && portfolio.UID != "" && uid != portfolio.UID {
		c.JSON(401, gin.H{"message": "許可がありません"})
		return
	}
	portfolio, err = putPortfolio(id, &p)
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
	portfolio, err := getPortfolio(id)
	if err != nil {
		if err.Error() == dbr.ErrNotFound.Error() {
			c.JSON(404, gin.H{"message": "データが見つかりませんでした"})
			return
		}
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	if uid, isExist := c.Get("AuthorizedUID"); isExist == false {
		c.JSON(500, gin.H{"message": "UID取得失敗"})
		return
	} else if portfolio.UID != "" && uid != portfolio.UID {
		c.JSON(401, gin.H{"message": "許可がありません"})
		return
	}
	err = deletePortfolio(id)
	if err != nil {
		c.JSON(500, gin.H{"message": "DB接続エラー"})
		return
	}
	c.Status(204)
}
