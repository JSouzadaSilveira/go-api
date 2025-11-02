package users

type Users struct {
	users []models.User
}


func New() *Users {
	return &Users{users: make([]models.User, 0)}
}

func (u *Users) GetAll() []models.User {
	return u.users
}

func (u *Users) GetByID(id uuid.UUID) *models.User {
	for _, user := range u.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (u Users) EmailExists(email string) bool {
	for _, user := range u.users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (u *Users) GetByEmail(email string) (*models.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}

func (u *Users) Add(user models.User) {
	u.users = append(u.users, user)
}