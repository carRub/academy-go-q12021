package controller

import (
	"fmt"
	"strconv"
	"net/http"
	
	"github.com/carRub/academy-go-q12021/model"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
)

// CharacterController defines what a controller must contain
type CharacterController struct {
	useCase CharacterUseCase
	render *render.Render
}

// CharacterUseCase defines functions to refer to our use cases
type CharacterUseCase interface {
	GetCharacters() ([]model.Character, error)
	GetCharacterByID(id int) (*model.Character, error) 
}

// NewCharacterController Creates a new character controller 
func NewCharacterController (cu CharacterUseCase, r *render.Render) CharacterController {
	c := CharacterController{
		useCase: cu,
		render: r,
	}

	return c
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
