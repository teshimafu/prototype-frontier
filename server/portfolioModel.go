package server

import (
	"time"

	dbr "github.com/gocraft/dbr/v2"
)

//PortfolioList is typeof PortfolioList information json
type PortfolioList struct {
	Data []Portfolio `json:"data"`
}

//Portfolio is typeof Portfolio detail json
type Portfolio struct {
	ID          int64         `db:"id" json:"id"`
	Title       string        `db:"title" json:"title"`
	UID         string        `db:"uid" json:"uid"`
	Author      string        `db:"author" json:"author"`
	Abstruct    string        `db:"abstruct" json:"abstruct"`
	Readme      string        `db:"readme" json:"readme"`
	Source      string        `db:"source" json:"source"`
	Link        string        `db:"link" json:"link"`
	AccessCount dbr.NullInt64 `db:"access_count" json:"access_count"`
	CreatedAt   time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at" json:"updated_at"`
}

//SearchQuery is query to search portfolio from DB
type SearchQuery map[string]string
