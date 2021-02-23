package main

import (
	_ "github.com/go-sql-driver/mysql"
	"local.packages/server"
)

func main() {
	server.Server()
}
