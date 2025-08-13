package controller

import (
	"encoding/json"
	"fmt"
	"main/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Personalidades)
}
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, p := range model.Personalidades {
		idInt, err := strconv.Atoi(id)

		if err != nil {
			json.NewEncoder(w).Encode(err)
		}

		if i+1 == idInt {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"err": "nada encontrado"})
}
