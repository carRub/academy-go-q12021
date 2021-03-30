package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CharacterController interface {
	GetCharacters(w http.ResponseWriter, r *http.Request)
	GetCharacterByID(w http.ResponseWriter, r *http.Request)
}

func NewRouter(c CharacterController) (http.Handler, error) {
	r := mux.NewRouter()

	r.HandleFunc("/characters", c.GetCharacters).Method(http.MethodGet).Name("GetAllCharacters")
	r.HandleFunc("/character/{id}", c.GetCharacterByID).Method(http.MethodGet).Name("GetCharacter")

	return r, nil
}
