package repository

import (
	"database/sql"

	"github.com/korasdor/go-chess/internal/domain"
)

type Users interface {
	Create(userData domain.UserData) error
	GetByCredentials(sidnInData domain.SignInData) (domain.UserData, error)
	GetById(userId int) (domain.UserData, error)
	UpdateUser(user domain.UserData) (domain.UserData, error)
}

type Repositories struct {
	UsersRepo Users
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UsersRepo: NewUsersRepo(db),
	}
}
