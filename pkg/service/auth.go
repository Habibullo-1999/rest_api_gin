package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/Habibullo-1999/rest_api_gin/pkg/repository"

)

const salt = "hjdsfds12034sfsdfsdf"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest_api_gin.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
