package server

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func getPortfolioList(searchOptions SearchQuery) ([](Portfolio), error) {
	sess, err := gocraftConnection()
	if err != nil {
		return []Portfolio{}, err
	}

	portfolioList := []Portfolio{}
	stmt := sess.Select("*").From("portfolio")
	if searchSQL := searchOptions.getWhereSQL(); searchSQL != "" {
		stmt = stmt.Where(searchSQL)
	}
	stmt = stmt.OrderDesc("id")
	stmt.Load(&portfolioList)
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

func getSearchQuery(c *gin.Context) SearchQuery {
	searchKeys := []string{"title", "author", "readme"}
	queries := make(SearchQuery)
	for _, v := range searchKeys {
		value := c.Query(v)
		if value != "" {
			value = strings.Replace(value, "'", "\"", -1)
			queries[v] = value
		}
	}
	return queries
}

func (sq *SearchQuery) getWhereSQL() string {
	sqlString := ""
	for k, v := range *sq {
		if len(sqlString) != 0 {
			sqlString += " OR "
		}
		sqlString += "portfolio." + k + " LIKE '%" + v + "%'"
	}
	return sqlString
}
