package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/carRub/academy-go-q12021/model"
)

type Service struct {
	fr *csv.Reader
	fw *csv.Writer
}

func NewCharacterService (fr *csv.Reader, fw *csv.Writer) (*Service, error) {
	// TODO: Implement?
	return &Service{fr, fw}, nil
}

func (s *Service) GetCharacters() ([]model.Character, error) {
	// TODO: Implement
	r := s.fr

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

func (s *Service) GetCharacterByID(id int) (*model.Character, error) {
	// TODO: Implementa
	r := s.fr
	fmt.Println(id)

	var character model.Character

	for {
		record, err := r.Read()
		
		if err == io.EOF {
			break
		}

		charID, _ := strconv.ParseInt(record[0], 10, 64)
		fileId := int(charID)

		if fileId == id {
			fmt.Println("character found")
			character.ID = fileId
			character.Name = record[1]
			character.Status = record[2]
			character.Species = record[3]
			character.Gender = record[4]
			break
		}
	}

	return &character, nil
}
