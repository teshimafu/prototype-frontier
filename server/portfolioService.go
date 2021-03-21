package server

import (
	"time"
)

func getPortfolio(ID int) (Portfolio, error) {
	sess, err := gocraftConnection()
	if err != nil {
		return Portfolio{}, err
	}

	portfolio := Portfolio{}
	err = sess.Select("*").From("portfolio").Where("id = ?", ID).LoadOne(&portfolio)
	return portfolio, err
}

func getPortfolioList() ([](Portfolio), error) {
	sess, err := gocraftConnection()
	if err != nil {
		return []Portfolio{}, err
	}

	var portfolioList []Portfolio
	sess.Select("*").From("portfolio").OrderDesc("id").Load(&portfolioList)
	return portfolioList, nil
}

func postPortfolio(p *Portfolio) (Portfolio, error) {
	sess, err := gocraftConnection()
	if err != nil {
		return Portfolio{}, err
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt

	err = sess.InsertBySql(`INSERT INTO "portfolio" ("title", "uid", "author", "abstruct", "readme", "source", "link", "created_at", "updated_at") 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING "id"`,
		p.Title, p.UID, p.Author, p.Abstruct, p.Readme, p.Source, p.Link, p.CreatedAt, p.UpdatedAt).Load(&p.ID)
	if err != nil {
		return Portfolio{}, err
	}
	return *p, nil
}

func putPortfolio(id int, p *Portfolio) (Portfolio, error) {
	sess, err := gocraftConnection()
	if err != nil {
		return Portfolio{}, err
	}
	portfolioMap := map[string]interface{}{
		"title":      p.Title,
		"author":     p.Author,
		"abstruct":   p.Abstruct,
		"readme":     p.Readme,
		"source":     p.Source,
		"link":       p.Link,
		"updated_at": time.Now(),
	}
	_, err = sess.Update("portfolio").
		SetMap(portfolioMap).
		Where("id = ?", id).Exec()
	if err != nil {
		return Portfolio{}, err
	}
	return getPortfolio(id)
}

func deletePortfolio(id int) error {
	sess, err := gocraftConnection()
	if err != nil {
		return err
	}
	_, err = sess.DeleteFrom("portfolio").
		Where("id = ?", id).Exec()
	return err
}
