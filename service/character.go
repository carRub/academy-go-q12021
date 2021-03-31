package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func (s *Service) InsertExternalCharacter(id int) error {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%v", id)
	w := s.fw

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var c model.Character
	if err := json.Unmarshal(data, &c); err != nil {
		log.Fatal("Service: Could not Unmarshal response", err)
	}

	var record []string
	record = append(record, strconv.FormatInt(int64(c.ID), 10), c.Name, c.Status, c.Species, c.Gender)
	if err := w.Write(record); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return nil
}
