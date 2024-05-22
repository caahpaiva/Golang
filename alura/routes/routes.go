package routes

import (
	"net/http"

	"github.com/caahpaiva/alura/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)

}
