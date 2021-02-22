package service

import (
	"local.packages/model"
)

//GetPortfolioList is return Portfoliolist
func GetPortfolioList() [](model.Portfolio) {
	db := gormConnect()
	defer db.Close()

	portfolioList := [](model.Portfolio){}

	db.Find(&portfolioList)
	return portfolioList
}
