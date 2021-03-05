package controller

import (
	"os"
	"encoding/csv"
	"strconv"
	"io"
	
	"github.com/carRub/academy-go-q12021/model"
	"github.com/carRub/academy-go-q12021/usecase"
)

// CharacterController defines what a controller must contain
type CharacterController struct {
	characterInteractor usecase.CharacterInteractor
}

// CharacterUseCase defines functions to refer to our use cases
type CharacterUseCase interface {
	GetCharacters() ([]model.Character)
	GetCharacterByID(id int) (*model.Character) 
}

// NewCharacterController Creates a new character controller 
func NewCharacterController (ci usecase.CharacterInteractor) CharacterController {
	return CharacterController{ci}
}

func (c CharacterController) GetCharacters() ([]model.Character, error) {
	csvFile, _ := os.Open("assets/characters.csv")
	defer csvFile.Close()
	r := csv.NewReader(csvFile)

	var character model.Character
	var characters []model.Character

	for {
		record, err := r.Read()
		
		if err == io.EOF {
			break
		}

		charID, _ := strconv.ParseInt(record[0], 10, 64)

		character.ID = int(charID)
		character.Name = record[1]
		character.Status = record[2]
		character.Species = record[3]
		character.Gender = record[4]

		characters = append(characters, character)
	}

	return characters, nil
}

func (c CharacterController) GetCharacterByID(id int) (*model.Character, error) {
	csvFile, _ := os.Open("assets/characters.csv")
	defer csvFile.Close()
	r := csv.NewReader(csvFile)

	var character model.Character

	for {
		record, err := r.Read()
		
		if err == io.EOF {
			break
		}

		charID, _ := strconv.ParseInt(record[0], 10, 64)
		if int(charID) == id {
			character.ID = int(charID)
			character.Name = record[1]
			character.Status = record[2]
			character.Species = record[3]
			character.Gender = record[4]
		}

	}

	return &character, nil
}