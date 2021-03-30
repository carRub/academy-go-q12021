package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/carRub/academy-go-q12021/controller"
	"github.com/carRub/academy-go-q12021/router"
	"github.com/carRub/academy-go-q12021/service"
	"github.com/carRub/academy-go-q12021/usecase"
	"github.com/unrolled/render"
)

func main() {
	csvFile, _ := os.Open("assets/characters.csv")
	fr := csv.NewReader(csvFile)
	fw := csv.NewWriter(csvFile)

	s, _ := service.NewCharacterService(fr, fw)
	i, _ := usecase.NewCharacterInteractor(s)
	c := controller.NewCharacterController(i, render.New())
	r, _ := router.NewRouter(c)

	fmt.Println("All layers created successfully", r)
}