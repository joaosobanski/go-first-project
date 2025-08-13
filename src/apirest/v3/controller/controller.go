package controller

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/model"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func CriarUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaP model.Personalidade
	json.NewDecoder(r.Body).Decode(&novaP)

	database.DB.Create(&novaP)
	json.NewEncoder(w).Encode(novaP)
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []model.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func AlteraUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p model.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p model.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

	// for i, p := range model.Personalidades {
	// 	idInt, err := strconv.Atoi(id)

	// 	if err != nil {
	// 		json.NewEncoder(w).Encode(err)
	// 	}

	// 	if i+1 == idInt {
	// 		json.NewEncoder(w).Encode(p)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(map[string]string{"err": "nada encontrado"})
}
