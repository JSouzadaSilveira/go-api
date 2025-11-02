package handlers

import (
	"github.com/jonathankossy/go-api/cmd/internal/usecases"
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
	http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	)
}