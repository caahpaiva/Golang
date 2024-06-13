package main

import (
	"github.com/caahpaiva/api-go-gin/database"
	"github.com/caahpaiva/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
