package repository

import (
	"database/sql"

	"github.com/korasdor/go-chess/internal/domain"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(userData domain.UserData) error {
	return nil
}

func (r *UsersRepo) GetByCredentials(sidnInData domain.SignInData) (domain.UserData, error) {

	return domain.UserData{}, nil
}

func (r *UsersRepo) GetById(userId int) (domain.UserData, error) {

	return domain.UserData{}, nil
}

func (r *UsersRepo) UpdateUser(user domain.UserData) (domain.UserData, error) {

	return domain.UserData{}, nil
}
