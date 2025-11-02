package repositories

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
		User: users.New(),
	}
}