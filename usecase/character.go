package usecase

import "github.com/carRub/academy-go-q12021/model"

// CharacterService interface to define the use cases applicable to this application/project
type CharacterService interface {
	GetCharacters() ([]model.Character, error)
	GetCharacterByID(id int) (*model.Character, error)
}

// CharacterInteractor defines the use case fields
type CharacterInteractor struct {
	Service CharacterService
}

// NewCharacterInteractor generates a new character usecase
func NewCharacterInteractor (s CharacterService) (*CharacterInteractor, error) {
	return &CharacterInteractor{s}, nil
}

// GetCharacters returns all the existing characters
func (c CharacterInteractor) GetCharacters() ([]model.Character, error) {
	characters, err := c.Service.GetCharacters()

	if err != nil {
		return nil, err
	}

	return characters, nil
}

// GetCharacterByID returns the character corresponding to a given Id
func (c CharacterInteractor) GetCharacterByID(id int) (*model.Character, error) {
	character, err := c.Service.GetCharacterByID(id)

	if err != nil {
		return nil, err
	}

	return character, nil
}
