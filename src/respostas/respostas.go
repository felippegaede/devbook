package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma repostas em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Erro retorna um erro em formato JSON para a requisição
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: erro.Error(),
	})
}
