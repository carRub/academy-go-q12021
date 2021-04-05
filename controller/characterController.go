package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/carRub/academy-go-q12021/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// CharacterController defines what a controller must contain
type CharacterController struct {
	useCase CharacterUseCase
	render  *render.Render
}

// CharacterUseCase defines functions to refer to our use cases
type CharacterUseCase interface {
	GetCharacters() ([]model.Character, error)
	GetCharacterByID(id int) (*model.Character, error)
	InsertExternalCharacter(id int) (*model.Character, error)
	GetCharactersConcurrently(t string, items int, itemsPerWorker int) ([]model.Character, error)
}

// NewCharacterController Creates a new character controller
func NewCharacterController(cu CharacterUseCase, r *render.Render) (CharacterController, error) {
	if cu == nil {
		return CharacterController{}, fmt.Errorf("Empty usecase")
	}

	if r == nil {
		return CharacterController{}, fmt.Errorf("Empty render")
	}

	c := CharacterController{
		useCase: cu,
		render:  r,
	}

	return c, nil
}

func (c CharacterController) GetCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := c.useCase.GetCharacters()
	if err != nil {
		err = fmt.Errorf("usecase request failed: %w", err)
		c.render.Text(w, http.StatusBadRequest, err.Error())

		return
	}

	c.render.JSON(w, http.StatusOK, characters)
}

func (c CharacterController) GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	reqId := mux.Vars(r)["id"]
	if reqId == "" {
		c.render.Text(w, http.StatusBadRequest, "Controller: param {id} must not be null")
		return
	}

	id, err := strconv.Atoi(reqId)
	if err != nil {
		c.render.Text(w, http.StatusBadRequest, "Controller: param {id} must be an integer")
		return
	}

	character, err := c.useCase.GetCharacterByID(id)
	if err != nil {
		err = fmt.Errorf("Usecase request failed %w", err)
		c.render.Text(w, http.StatusBadRequest, err.Error())

		return
	}

	c.render.JSON(w, http.StatusOK, character)
}

func (c CharacterController) InsertExternalCharacter(w http.ResponseWriter, r *http.Request) {
	reqId := mux.Vars(r)["id"]
	if reqId == "" {
		c.render.Text(w, http.StatusBadRequest, "Controller: param {id} must not be null")
		return
	}

	id, err := strconv.Atoi(reqId)
	if err != nil {
		c.render.Text(w, http.StatusBadRequest, "Controller: param {id} must be an integer")
		return
	}

	character, err := c.useCase.InsertExternalCharacter(id)
	if err != nil {
		err = fmt.Errorf("Usecase request failed %w", err)
		c.render.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	c.render.JSON(w, http.StatusOK, character)
}

func (c CharacterController) GetCharactersConcurrently(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()
	reqType := u.Get("type")
	reqItems := u.Get("items")
	reqItemsPerWorkers := u.Get("items_per_workers")
	/*reqType := mux.Vars(r)["type"]
	reqItems := mux.Vars(r)["items"]
	reqItemsPerWorker := mux.Vars(r)["items_per_worker"]*/

	if reqType == "" {
		c.render.Text(w, http.StatusBadRequest, "Controller: Type must not be null")
	}

	if reqItems == "" {
		c.render.Text(w, http.StatusBadRequest, "Controller: Type must not be null")
	}

	if reqItemsPerWorkers == "" {
		c.render.Text(w, http.StatusBadRequest, "Controller: Type must not be null")
	}

	charItems, err := strconv.ParseInt(reqItems, 10, 64)
	if err != nil {
		log.Fatal("Controller: Could not cast string to int32 - items")	
	}
	items := int(charItems)

	charItemsPerWorker, err := strconv.ParseInt(reqItemsPerWorkers, 10, 64)
	if err != nil {
		log.Fatal("Controller: Could not cast string to int32 - items_per_worker")	
	}
	itemsPerWorker := int(charItemsPerWorker)

	characters, err := c.useCase.GetCharactersConcurrently(reqType, items, itemsPerWorker)

	c.render.JSON(w, http.StatusOK, characters)
}