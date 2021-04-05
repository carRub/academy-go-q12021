package main

import (
	"log"
	"net/http"

	"github.com/carRub/academy-go-q12021/controller"
	"github.com/carRub/academy-go-q12021/router"
	"github.com/carRub/academy-go-q12021/service"
	"github.com/carRub/academy-go-q12021/usecase"

	"github.com/unrolled/render"
)

func main() {
	url := "https://rickandmortyapi.com/api/character/"
	file := "assets/characters.csv"

	s, err := service.NewCharacterService(url, file)
	if err != nil {
		log.Fatal("could not instance service")
	}

	i, err := usecase.NewCharacterInteractor(s)
	if err != nil {
		log.Fatal("could not instance useCase")
	}

	c, err := controller.NewCharacterController(i, render.New())
	if err != nil {
		log.Fatal("could not instance controller")
	}

	r, err := router.NewRouter(c)
	if err != nil {
		log.Fatal("could not instance router")
	}

	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
