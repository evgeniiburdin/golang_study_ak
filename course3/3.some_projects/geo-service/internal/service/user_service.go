package service

import (
	"errors"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"

	"geo-service/models"
)

type UserServicer interface {
	Login(user models.User) (string, error)
	Register(user models.User) error
}

type UserService struct {
	repo      UserRepository
	tokenAuth *jwtauth.JWTAuth
}

type UserRepository interface {
	ReadUser(username string) (string, error)
	CreateUser(username, password string) error
}

func NewUserService(repo UserRepository, tokenAuth *jwtauth.JWTAuth) *UserService {
	return &UserService{repo: repo, tokenAuth: tokenAuth}
}

func (s *UserService) Login(user models.User) (string, error) {
	storedPassword, err := s.repo.ReadUser(user.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	_, tokenString, _ := s.tokenAuth.Encode(map[string]interface{}{"username": user.Username})
	return tokenString, nil
}

func (s *UserService) Register(user models.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(user.Username, string(encryptedPassword))
}
