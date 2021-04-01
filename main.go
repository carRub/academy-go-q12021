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

	s, _ := service.NewCharacterService(url, file)
	i, _ := usecase.NewCharacterInteractor(s)
	c, _ := controller.NewCharacterController(i, render.New())
	r, _ := router.NewRouter(c)

	log.Fatal(http.ListenAndServe(":3000", r))
}
