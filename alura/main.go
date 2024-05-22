package main

import (
	"net/http"

	"github.com/caahpaiva/alura/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":800", nil)
}
