package main

import (
	"github.com/jonathankossy/go-api/cmd/internal/handlers"
	"github.com/jonathankossy/go-api/cmd/internal/repositories"
	"github.com/jonathankossy/go-api/cmd/internal/usecases"
)

// Cadastrar e listar usuários
// handler <- usecases (service/domain) <- repositories <- models
func main() {
	repositories := repositories.New()
	useCases := usecases.New(repositories)
	handlers := handlers.New(useCases)

	handlers.Listen(8080)
}