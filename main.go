package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"local.packages/server"
)

func main() {
	server.Server()
}
