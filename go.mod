module gin_mod

// +heroku goVersion go1.15
go 1.15

replace local.packages/server => ./server

require (
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.9.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	local.packages/server v0.0.0-00010101000000-000000000000
)
