package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/paulormbsd/zeus-hammer/models"
)

// Função que cria a home page ou página para utilização do módulo http, se atentar a questão do atalho WR para utilizar o ResponseWriter e Request
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

// Função que retorna todas as personalidades mocadas ou em database para a página, o encode ajuda a mandar a request em formato json
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

// Função que retorna um valor, baseado em ID
func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]
	var personalidade models.Personalidade
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]
	var personalidade models.Personalidade
	//database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	//database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
