package services

import (
	"time"

	"github.com/korasdor/go-chess/internal/domain"
	"github.com/korasdor/go-chess/internal/repository"
	"github.com/korasdor/go-commons/auth"
	"github.com/korasdor/go-commons/cache"
	"github.com/korasdor/go-commons/hash"
)

type Authorization interface {
	SignUp(signUpData domain.SignUpData) error
	SignIn(signInData domain.SignInData) (domain.Tokens, error)
	RefreshTokens(refreshToken string) (domain.Tokens, error)
}

type Users interface {
	GetUser(userId int) (domain.UserData, error)
	UpdateUser(user domain.UserData) (domain.UserData, error)
	DeleteUser(userId int) error
}

type Services struct {
	AuthorizationService Authorization
	UsersService         Users
}

type Deps struct {
	Repos           *repository.Repositories
	Cache           cache.Cache
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps *Deps) *Services {
	return &Services{
		AuthorizationService: NewAuthorizationService(deps.Repos, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL),
		UsersService:         NewUsersService(deps.Repos),
	}
}
