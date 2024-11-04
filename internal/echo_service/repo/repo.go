package repo

import "example/internal/common/models"

type Repo struct {
}

// TODO: make repo
func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetUser(userID string) (*models.User, error) {

	return &models.User{
		ID:   userID,
		Name: "Test User",
	}, nil
}
