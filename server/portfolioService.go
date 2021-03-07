package server

import (
	"time"
)

func getPortfolio(ID int) (Portfolio, error) {
	db, err := gormConnect()
	if err != nil {
		return Portfolio{}, err
	}
	defer db.Close()

	portfolio := Portfolio{}

	err = db.Find(&portfolio, ID).Error
	if err != nil {
		return Portfolio{}, err
	}
	return portfolio, nil
}

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

func postPortfolio(p *Portfolio) (Portfolio, error) {

	db, err := gormConnect()
	if err != nil {
		return Portfolio{}, err
	}
	defer db.Close()

	p.CreatedAt = time.Now()
	db.Create(p)
	return *p, nil
}

func putPortfolio(p *Portfolio) (Portfolio, error) {

	db, err := gormConnect()
	if err != nil {
		return Portfolio{}, err
	}
	defer db.Close()

	p.UpdatedAt = time.Now()
	db.Save(p)
	return *p, nil
}
