package main

import (
	"api/internal/handlers"
	"api/internal/repositories"
	"api/internal/usecases"
)
// Cadastrar e listar usuários
// handler <- usecases (service/domain) <- repositories <- models
func main() {
	repositories := repositories.New()
	useCases := usecases.New(repositories)
	handlers := handlers.New(useCases)

	handlers.Listen(8080)
}