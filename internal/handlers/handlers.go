package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"api/internal/usecases"
)

type Handlers struct {
	UseCases *usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{
		UseCases: useCases,
	}
}

func (h Handlers) Listen(port int) error {
	h.registerUserEndpoints()

	slog.Info("Starting server on port %d", port)
	
	return http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	)
}