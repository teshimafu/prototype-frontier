package controller

import (
	"github.com/gin-gonic/gin"
	"local.packages/model"
)

//GetPortfolioList is return Portfoliolist
func GetPortfolioList(c *gin.Context) {
	portfolioList := model.PortfolioList{
		Data: []model.Portfolio{
			{
				ID:         1,
				Title:      "サーバーのデータ",
				Author:     "teshima",
				CreateDate: "1/1",
				Tags:       []string{"1", "2", "3"},
				Abstruct:   "## title etc...",
				Source:     "https:google.com",
				URL:        "https://qiita.com",
			},
		},
	}
	c.JSON(200, portfolioList)
}
