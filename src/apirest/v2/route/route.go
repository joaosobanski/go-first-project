package route

import (
	"log"
	"main/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/personalidades", controller.TodasPersonalidades)
	r.HandleFunc("/personalidades/{id}", controller.RetornaUmaPersonalidade)
	log.Fatal(http.ListenAndServe(":8000", r))
}
