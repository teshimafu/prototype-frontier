package main

import (
	_ "github.com/lib/pq"
	"local.packages/server"
)

func main() {
	server.Server()
}
