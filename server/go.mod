module server

go 1.15

replace local.packages/controller => ./controller

require (
	github.com/gin-contrib/static v0.0.0-20200916080430-d45d9a37d28e
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.6.3
	local.packages/controller v0.0.0-00010101000000-000000000000
)
