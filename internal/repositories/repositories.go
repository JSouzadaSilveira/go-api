package repositories

import (
	"api/internal/models"
	"github.com/google/uuid"
)

type Repositories struct {
	User interface{
		GetAll() []models.User
		Add(user models.User)
		GetByID(id uuid.UUID) *models.User
		EmailExists(email string) bool
		GetByEmail(email string) (*models.User, error)
	}
}

func New() *Repositories {
	return &Repositories{
		User: NewUsers(),
	}
}