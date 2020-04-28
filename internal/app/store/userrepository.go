package store

import "gitlab.devkeeper.com/authreg/server/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindBylogin(login string) (*model.User, error) {
	return nil, nil
}
