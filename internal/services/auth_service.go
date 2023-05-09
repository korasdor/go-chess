package services

import (
	"fmt"
	"time"

	"github.com/korasdor/go-chess/internal/domain"
	"github.com/korasdor/go-chess/internal/repository"
	"github.com/korasdor/go-commons/auth"
	"github.com/korasdor/go-commons/hash"
)

type AuthorizationService struct {
	repo            *repository.Repositories
	hasher          hash.PasswordHasher
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthorizationService(
	repo *repository.Repositories,
	hasher hash.PasswordHasher,
	tokenManager auth.TokenManager,
	accessTokenTTL time.Duration,
	refeshTokenTTL time.Duration,
) *AuthorizationService {

	return &AuthorizationService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refeshTokenTTL,
	}
}

func (s *AuthorizationService) SignUp(signUpData domain.SignUpData) error {
	passwordHash, err := s.hasher.Hash(signUpData.Password)
	if err != nil {
		return err
	}

	userData := domain.UserData{
		Username:     signUpData.Name,
		Password:     passwordHash,
		Phone:        signUpData.Phone,
		Email:        signUpData.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
	}

	return s.repo.UsersRepo.Create(userData)
}

func (s *AuthorizationService) SignIn(signInData domain.SignInData) (domain.Tokens, error) {
	var token domain.Tokens

	passwordHash, err := s.hasher.Hash(signInData.Password)
	if err != nil {
		return token, err
	}

	signInData.Password = passwordHash
	user, err := s.repo.UsersRepo.GetByCredentials(signInData)
	if err != nil {
		return token, err
	}

	token, err = s.generateTokens(fmt.Sprintf(`%d`, user.Id))
	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthorizationService) RefreshTokens(refreshToken string) (domain.Tokens, error) {
	var token domain.Tokens

	userId, err := s.tokenManager.ParseJWT(refreshToken)
	if err != nil {
		return token, err
	}

	token, err = s.generateTokens(userId)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthorizationService) generateTokens(userId string) (domain.Tokens, error) {
	var token domain.Tokens

	accessToken, err := s.tokenManager.NewJWT(userId, s.accessTokenTTL)
	if err != nil {
		return token, err
	}

	refreshToken, err := s.tokenManager.NewJWT(userId, s.refreshTokenTTL)
	if err != nil {
		return token, err
	}

	token = domain.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
