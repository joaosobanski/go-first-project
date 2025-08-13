package route

import (
	"log"
	"main/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("Get")
	r.HandleFunc("/personalidades", controller.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controller.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controller.AlteraUmaPersonalidade).Methods("Put")
	r.HandleFunc("/personalidades", controller.CriarUmaNovaPersonalidade).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
