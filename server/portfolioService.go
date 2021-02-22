package server

//getPortfolioList is return Portfoliolist
func getPortfolioList() [](Portfolio) {
	db := gormConnect()
	defer db.Close()

	portfolioList := [](Portfolio){}

	db.Find(&portfolioList)
	return portfolioList
}
