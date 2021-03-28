package server

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocraft/dbr/v2"
)

func getPortfolio(id int) (Portfolio, error) {
	sess, err := gocraftConnection()
	if err != nil {
		return Portfolio{}, err
	}

	p := Portfolio{}
	err = sess.Select("*").From("portfolio").Where("id = ?", id).LoadOne(&p)
	if err != nil {
		return Portfolio{}, err
	}
	err = countUpPortfolio(sess, &p)
	return p, err
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

func countUpPortfolio(sess *dbr.Session, p *Portfolio) error {
	if p.AccessCount.Valid {
		p.AccessCount = dbr.NewNullInt64(p.AccessCount.Int64 + 1)
	} else {
		p.AccessCount = dbr.NewNullInt64(1)
	}
	portfolioMap := map[string]interface{}{
		"title":        p.Title,
		"author":       p.Author,
		"abstruct":     p.Abstruct,
		"readme":       p.Readme,
		"source":       p.Source,
		"link":         p.Link,
		"access_count": p.AccessCount,
		"updated_at":   time.Now(),
	}
	_, err := sess.Update("portfolio").
		SetMap(portfolioMap).
		Where("id = ?", p.ID).Exec()
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
