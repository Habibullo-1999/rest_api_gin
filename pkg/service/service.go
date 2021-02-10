package service

import (
	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/Habibullo-1999/rest_api_gin/pkg/repository"

)

type Authorization interface {
	CreateUser(user rest_api_gin.User) (int, error)
	CreateToToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type NoteList interface {
}

type NoteItem interface {
}

type Service struct {
	Authorization
	NoteList
	NoteItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
