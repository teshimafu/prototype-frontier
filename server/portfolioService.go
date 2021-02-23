package server

import "time"

func getPortfolioList() ([](Portfolio), error) {
	db, err := gormConnect()
	if err != nil {
		return [](Portfolio){}, err
	}
	defer db.Close()

	portfolioList := [](Portfolio){}

	db.Find(&portfolioList)
	return portfolioList, nil
}

func postPortfolio(p *Portfolio) error {
	db, err := gormConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	p.CreatedAt = time.Now()
	db.Create(p)
	return nil
}
