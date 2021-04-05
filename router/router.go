package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CharacterController interface {
	GetCharacters(w http.ResponseWriter, r *http.Request)
	GetCharacterByID(w http.ResponseWriter, r *http.Request)
	InsertExternalCharacter(w http.ResponseWriter, r *http.Request)
	GetCharactersConcurrently(w http.ResponseWriter, r *http.Request)
}

func NewRouter(c CharacterController) (http.Handler, error) {
	r := mux.NewRouter()

	r.HandleFunc("/characters", c.GetCharacters).Methods(http.MethodGet).Name("GetAllCharacters")
	r.HandleFunc("/character/{id}", c.GetCharacterByID).Methods(http.MethodGet).Name("GetCharacter")
	r.HandleFunc("/character/external/{id}", c.InsertExternalCharacter).Methods(http.MethodGet).Name("GetExternalCharacter")
	r.HandleFunc("/character/concurrent/", c.GetCharactersConcurrently).Methods(http.MethodGet).Name("GetCharactersConcurrently")

	return r, nil
}
