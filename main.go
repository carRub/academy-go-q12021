package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/carRub/academy-go-q12021/controller"
	"github.com/carRub/academy-go-q12021/router"
	"github.com/carRub/academy-go-q12021/service"
	"github.com/carRub/academy-go-q12021/usecase"
	"github.com/unrolled/render"
)

func main() {
	csvFile, err := os.Open("assets/characters.csv")
	if err != nil {
		log.Fatal("Error creating file reader", err)
	}
	fr := csv.NewReader(csvFile)

	csvW, err := os.OpenFile("assets/characters.csv", os.O_APPEND|os.O_WRONLY,  os.ModeAppend)
	if err != nil {
		log.Fatal("Error creating file writer")
	}
	fw := csv.NewWriter(csvW)
	
	defer csvFile.Close()
	defer csvW.Close()

	s, _ := service.NewCharacterService(fr, fw)
	i, _ := usecase.NewCharacterInteractor(s)
	c := controller.NewCharacterController(i, render.New())
	r, _ := router.NewRouter(c)

	log.Fatal(http.ListenAndServe(":3000", r))
}
