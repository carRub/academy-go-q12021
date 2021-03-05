package main

import (

	"fmt"

	"github.com/carRub/academy-go-q12021/controller"
	"github.com/carRub/academy-go-q12021/usecase"
)

func main() {
	var characterService usecase.CharacterService
	characterInteractor, _ := usecase.NewCharacterInteractor(characterService)
	characterController := controller.NewCharacterController(*characterInteractor)

	fmt.Println(characterController)

}