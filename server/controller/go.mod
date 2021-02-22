module controller

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	local.packages/model v0.0.0-00010101000000-000000000000
)

replace local.packages/model => ../model
