package usecases

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u *UseCases) GetAll() []models.User {
	users := u.repos.User.GetAll()
	return users
}

func (u *UseCases) AddUser(user models.User) (uuid.UUID, error) {
	exists := u.repos.User.EmailExists(user.Email)
	if exists {
		slog.Info("email already in use", "email", user.Email)
		return uuid.Nil, fmt.Errorf("email %s already in use", user.Email)
	}
	repoReq := models.User{
		ID: uuid.New(),
		Name: user.Name,
		Email: user.Email,
	}

	u.repos.User.Add(repoReq)
	return repoReq.ID, nil
}