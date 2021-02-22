package model

//PortfolioList is typeof PortfolioList information json
type PortfolioList struct {
	Data []Portfolio `json:"data"`
}

//Portfolio is typeof Portfolio detail json
type Portfolio struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Title    string `gorm:"type:varchar(256)" json:"title"`
	Author   string `gorm:"type:varchar(256)" json:"author"`
	Abstruct string `gorm:"type:text" json:"abstruct"`
	Source   string `gorm:"type:varchar(1024)" json:"source"`
	Link     string `gorm:"type:varchar(1024)" json:"link"`
}
