module gin_mod

go 1.15

replace local.packages/server => ./server

replace local.packages/controller => ./server/controller

replace local.packages/model => ./server/model

require (
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-contrib/static v0.0.0-20200916080430-d45d9a37d28e // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/sys v0.0.0-20210220050731-9a76102bfb43 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	local.packages/controller v0.0.0-00010101000000-000000000000 // indirect
	local.packages/server v0.0.0-00010101000000-000000000000 // indirect
)
