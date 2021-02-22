package model

//PortfolioList is typeof PortfolioList information json
type PortfolioList struct {
	Data []Portfolio `json:"data"`
}

//Portfolio is typeof Portfolio detail json
type Portfolio struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	CreateDate string   `json:"createDate"`
	Tags       []string `json:"tags"`
	Abstruct   string   `json:"abstruct"`
	Source     string   `json:"source"`
	URL        string   `json:"url"`
}
