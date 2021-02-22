module service

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	local.packages/model v0.0.0-00010101000000-000000000000
)

replace local.packages/model => ../model
