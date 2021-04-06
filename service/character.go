package service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/carRub/academy-go-q12021/model"
)

type Service struct {
	url  string
	file string
}

func NewCharacterService(url, file string) (*Service, error) {
	if url == "" {
		return nil, fmt.Errorf("Empty url")
	}

	if file == "" {
		return nil, fmt.Errorf("Empty file path")
	}

	return &Service{url, file}, nil
}

func (s *Service) GetCharacters() ([]model.Character, error) {
	characters, err := readAllRecordsFromCsv(s)
	if err != nil {
		return nil, fmt.Errorf("Could not get characters %w", err)
	}

	return characters, nil
}

func (s *Service) GetCharacterByID(id int) (*model.Character, error) {
	character, err := readRecordFromCsv(s, id)
	if err == io.EOF {
		return nil, fmt.Errorf("Character not found in file %w", err)
	}

	return &character, nil
}

func (s *Service) InsertExternalCharacter(id int) (*model.Character, error) {
	reqUrl := s.url + strconv.Itoa(id)

	csvW, err := os.OpenFile(s.file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, fmt.Errorf("Error creating file writer %w", err)
	}
	w := csv.NewWriter(csvW)

	defer csvW.Close()

	res, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		res.Body.Close()
		return nil, fmt.Errorf("Could not read response body %w", err)
	}
	res.Body.Close()

	var c model.Character
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("Service: could not Umarshal response %w", err)
	}

	var record []string
	record = append(record, strconv.FormatInt(int64(c.ID), 10), c.Name, c.Status, c.Species, c.Gender)
	if err := w.Write(record); err != nil {
		return nil, fmt.Errorf("Writing to csv %w", err)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return nil, fmt.Errorf("writer error %w", err)
	}

	return &c, nil
}

func (s *Service) GetCharactersConcurrently(t string, items int, itemsPerWorkers int) ([]model.Character, error) {
	workers := int(math.Ceil((float64(items) / float64(itemsPerWorkers))))
	fmt.Println(workers)

	rows, err := readAllRecordsFromCsv(s)
	if err != nil {
		return nil, fmt.Errorf("Could not get characters %w", err)
	}

	results := make(chan model.Character, items)
	shutdown := make(chan struct{})
	jobs := make(chan int, len(rows))

	csvFile, err := os.Open(s.file)
	if err != nil {
		return nil, fmt.Errorf("Error creating file reader %w", err)
	}
	r := csv.NewReader(csvFile)
	defer csvFile.Close()

	for w := 1; w <= 3; w++ {
		go worker(r, jobs, shutdown, results)
	}

	for i := 1; i <= items; i++ {
		jobs <- i
	}
	close(jobs)

	var characters []model.Character
	for character := range results {
		fmt.Println("iterating Results")
		fmt.Println("Characters: ", characters)

		if isOfType(t, character.ID) {
			fmt.Println("appending")
			characters = append(characters, character)
		}

		if len(characters) == itemsPerWorkers {
			fmt.Println("leaving results for")
			break
		}
	}
	fmt.Println("shutting down")
	close(shutdown)

	return characters, nil
}

func isOfType(t string, id int) bool {
	if t == "odd" && id%2 != 0 {
		return true
	}

	if t == "even" && id%2 == 0 {
		return true
	}

	return false
}

func readRecordFromCsv(s *Service, id int) (model.Character, error) {
	csvFile, err := os.Open(s.file)
	if err != nil {
		log.Fatal("Error creating file reader", err)
	}
	r := csv.NewReader(csvFile)
	defer csvFile.Close()

	var character model.Character

	for {
		record, err := r.Read()

		if errors.Is(err, io.EOF) {
			return model.Character{}, err
		}

		charID, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Fatal("Service: couldn't cast string to int")
		}

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

	return character, nil
}

func readAllRecordsFromCsv(s *Service) ([]model.Character, error) {
	csvFile, err := os.Open(s.file)
	if err != nil {
		log.Fatal("Error creating file reader", err)
	}
	r := csv.NewReader(csvFile)
	defer csvFile.Close()

	var character model.Character
	var characters []model.Character

	for {
		record, err := r.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		charID, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			log.Fatal("Service/getCharacters: couldn't cast string to int")
		}

		character.ID = int(charID)
		character.Name = record[1]
		character.Status = record[2]
		character.Species = record[3]
		character.Gender = record[4]

		characters = append(characters, character)
	}

	return characters, nil
}

func worker(r *csv.Reader, jobs <-chan int, shutdown <-chan struct{}, results chan model.Character) {
	for job := range jobs {
		select {
		case <-shutdown:
			//log
			fmt.Println("shutting down from worker")
			return
		default:
			fmt.Println("The job is: ", job)
			record, err := r.Read()

			if errors.Is(err, io.EOF) {
				break
			}

			var character model.Character
			charID, err := strconv.ParseInt(record[0], 10, 64)
			if err != nil {
				log.Fatal("worker: couldn't cast string to int")
			}

			character.ID = int(charID)
			character.Name = record[1]
			character.Status = record[2]
			character.Species = record[3]
			character.Gender = record[4]
			
			fmt.Println(character)
			results <- character
		
		}
	}
	fmt.Println("leaving worker")
}
